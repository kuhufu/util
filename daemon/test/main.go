package main

import (
	"github.com/kuhufu/util/daemon"
	"os"
	"time"
)

func main() {
	daemon.Run(new(App))
}

type App struct {
}

func (a *App) Start() {
	daemon.Log(os.Args)

	var i int
	for {
		time.Sleep(time.Second)
		daemon.Log(i)
		i++

		if i == 10 {
			//panic("err")
		}
	}
}

func (a *App) Daemon() bool {
	return true
}

func (a *App) Name() string {
	return "test"
}
