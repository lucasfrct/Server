package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/schedule"
	"github.com/lucasfrct/servertools/tasks"
)

func main() {

	done := make(chan bool)
	var scheduleTime time.Duration = 10

	done = schedule.Schedule(func() {

		spew.Dump("chamada schedule")
		// tasks.Pull("../servertools")

		tasks.Push("../servertools")
		// asd

		// err := tasks.Copy("../sinis", "../transfer")
		// if err != nil {
		// 	fmt.Print("Error when copying!", err.Error())
		// 	panic(err)
		// }
		// fmt.Print("Copied successfully!")

		// tasks.Copy("../transfer", "../sinis")

		// // Maquina 1 - Fluxo In (M1Fi)
		// pathSourceM1Fi := "../transfering"
		// pathDestM1Fi := "../sinis"
		// cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
		// filesM1Fi := tasks.PathfinderOfFilesModifieds(pathSourceM1Fi)
		// filesM1FiCopied := tasks.CopyFiles(filesM1Fi, pathSourceM1Fi, pathDestM1Fi)
		// spew.Dump(cmdStrPullM1Fi)

		// // Maquina 1 - Fluxo out(M1Fo)
		// pathSourceM1Fo := "../sinis"
		// pathDestM1Fo := "../transfering"
		// filesM1Fo := tasks.PathfinderOfFilesModifieds(pathSourceM1Fo)
		// filesM1FoCopied := tasks.CopyFiles(filesM1Fo, pathSourceM1Fo, pathDestM1Fo)
		// resp := tasks.CommitFilesModified(pathDestM1Fo)
		// spew.Dump(filesM1FoCopied)

		// // Maquina 2 - Fluxo in(M2Fi)
		// pathSourceM2Fi := "../transfering"
		// cmdStrPullM2Fi := gitcommand.GitPull(pathSourceM2Fi)
		// spew.Dump(cmdStrPullM2Fi)

		// // Maquina 2 - Fluxo out(M2Fo)
		// pathSourceM2Fo := "../transfering"
		// respM2Fo := tasks.CommitFilesModified(pathSourceM2Fo)
		// spew.Dump(respM2Fo)

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
