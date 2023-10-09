package tasks

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
	"github.com/lucasfrct/servertools/pkg/modules/schedule"
)

func Pull(pathSourceM1Fi string) {

	done := make(chan bool)
	var scheduleTime time.Duration = 10

	done = schedule.Schedule(func() {

		spew.Dump("test")

		// Maquina 1 - Fluxo In (M1Fi)
		// pathSourceM1Fi := "../transfering"
		cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
		spew.Dump(cmdStrPullM1Fi)

	}, scheduleTime*time.Second)

	go func() {
		time.Sleep(480 * time.Minute)
		done <- true
	}()

	select {
	case v1 := <-done:
		fmt.Println("END: ", v1)
	}

}
