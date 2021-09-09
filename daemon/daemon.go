package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type App interface {
	Start()
	Daemon() bool
	Name() string
}

type exitCodeItem struct {
	ExitCode int
	Cmd      *exec.Cmd
}

var (
	notifyChan   chan os.Signal
	exitCodeChan chan *exitCodeItem
	curCmd       atomic.Value
	timeout      time.Time
)

func init() {
	exitCodeChan = make(chan *exitCodeItem, 1)

	timeout = time.Now().Add(time.Second * 3)

	notifyChan = make(chan os.Signal, 1)
	signal.Notify(notifyChan, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT)
}

func isChild() bool {
	return os.Getenv("is_kuhufu_child") != ""
}

func Run(app App) {
	if isChild() {
		Log("parent pid:", os.Getppid())
		app.Start()
		return
	}

	if app.Daemon() {
		Log("daemon")
		_, err := runWithDaemon()
		if err != nil {
			Log("first child start failed")
			return
		}

		go handleExitCode()
		go handleSignal()
	} else {
		app.Start()
	}

	select {}
}

func execInChild() (*exec.Cmd, error) {
	cmd := exec.Command(os.Args[0], os.Args[1:]...)

	cmd.Env = append(os.Environ(), "is_kuhufu_child=true") //传给子进程的环境参数

	cmd.Stderr = os.Stderr //将子进程的标准错误重定向到父进程的标准错误
	cmd.Stdout = os.Stdout

	curCmd.Store(cmd)

	if err := cmd.Start(); err != nil {
		Log("child err ", err)
		return cmd, err
	}

	Log("child pid:", cmd.Process.Pid)

	return cmd, nil
}

func runWithDaemon() (*exec.Cmd, error) {
	cmd, err := execInChild()
	if err != nil {
		return nil, err
	}

	go func() {
		ps, err := cmd.Process.Wait() //调用wait清理僵尸子进程
		if err != nil {
			Log("wait child err:", err)
		}

		Log("child state:", ps)
		Log("child exit code:", ps.ExitCode())

		exitCodeChan <- &exitCodeItem{
			ExitCode: ps.ExitCode(),
			Cmd:      cmd,
		}
	}()

	return cmd, nil
}

func handleExitCode() {
	for v := range exitCodeChan {
		if timeout.After(time.Now()) {
			Log("start failed so fast")
			os.Exit(0)
		}

		exitCode, cmd := v.ExitCode, v.Cmd

		Log("receive child exit code", exitCode)

		if exitCode == 0 || exitCode == -1 {
			Log(fmt.Sprintf("child %v normal exit with code %v, not restart", cmd.Process.Pid, exitCode))
		} else {
			Log(fmt.Sprintf("child %v abnormal exit with code %v, restart", cmd.Process.Pid, exitCode))
			_, err := runWithDaemon()
			if err != nil {
				Log("exit code restart failed")
			}
		}
	}
}

func handleSignal() {
	for {

		select {
		case sig := <-notifyChan:
			log.Println()
			Log("receive signal:", sig)

			switch sig {
			case syscall.SIGHUP:
				//检查可执行文件是否存在
				if !FileExist(os.Args[0]) {
					Log(os.Args[0], "not exist restart failed")
					continue
				}

				err := getCmd().Process.Kill()
				if err != nil {
					Log("kill child fail", err)
				}

				Log("重启")
				runWithDaemon()
			default:
				Log("准备退出")
				getCmd().Process.Kill()
				Log("kill child", getCmd().Process.Pid)
				Log("exit")
				os.Exit(0)
			}
		}
	}
}

func getCmd() *exec.Cmd {
	return curCmd.Load().(*exec.Cmd)
}
