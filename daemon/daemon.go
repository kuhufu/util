package daemon

import (
	"fmt"
	"github.com/kuhufu/util/concurrency"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type App interface {
	Start()
	Daemon() bool
	Name() string
}

var notifyChan chan os.Signal
var wg concurrency.WaitGroupWrapper

func init() {
	notifyChan = make(chan os.Signal, 1)
	signal.Notify(notifyChan, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT)
}

func IsChild() bool {
	return os.Getenv("is_kuhufu_child") != ""
}

func Run(app App) {
	if IsChild() {
		Log("is child, parent pid:", os.Getppid())
		app.Start()
		return
	}

	if app.Daemon() {
		Log("daemon")
		RunWithDaemon()
	} else {
		app.Start()
	}
}

func execInChild() *exec.Cmd {
	cmd := exec.Command(os.Args[0], os.Args[1:]...)

	cmd.Env = append(os.Environ(), "is_kuhufu_child=true") //传给子进程的环境参数

	cmd.Stderr = os.Stderr //将子进程的标准错误重定向到父进程的标准错误
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		Log("child err ", err)
		os.Exit(0)
	}

	Log("child process pid:", cmd.Process.Pid)

	return cmd
}

func RunWithDaemon() {
	var exitCodeChan chan int
	exitCodeChan = make(chan int)

	cmd := execInChild()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				Log(err)
			}
		}()
		ps, err := cmd.Process.Wait() //调用wait清理僵尸子进程
		if err != nil {
			Log("wait child err:", err)
		}
		Log("child state:", ps)
		Log("child exit code:", ps.ExitCode())

		exitCodeChan <- ps.ExitCode()
	}()

	for {
		select {
		case exitCode := <-exitCodeChan:
			Log("receive child exit code", exitCode)

			if exitCode == 0 || exitCode == -1 {
				Log(fmt.Sprintf("child %v normal exit with code %v, not restart", cmd.Process.Pid, exitCode))
			} else {
				Log(fmt.Sprintf("child %v abnormal exit with code %v, restart", cmd.Process.Pid, exitCode))
				RunWithDaemon()
				return
			}
		case sig := <-notifyChan:
			Log("receive signal:", sig)

			switch sig {
			case syscall.SIGHUP:
				//检查可执行文件是否存在
				_, err2 := os.Stat(os.Args[0])
				if os.IsNotExist(err2) {
					Log(os.Args[0], "not exist restart failed")
					continue
				}

				err := cmd.Process.Kill()
				if err != nil {
					Log("kill child fail", err)
				}

				Log("重启")
				RunWithDaemon()
				return
			default:
				cmd.Process.Kill()
				Log("退出")
				return
			}
		}
	}
}
