package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
	"github.com/lucasfrct/servertools/pkg/modules/schedule"
)

func main() {

	done := make(chan bool)
	var scheduleTime time.Duration = 60

	done = schedule.Schedule(func() {

		// pathSource := "../utils"
		// pathDest := "../cp"
		// files := TaskPathfinderOfFilesModifieds(pathSource)
		// _ = TaskCopyFiles(files, pathSource, pathDest)
		// resp := TaskCommitFilesModified(pathDest)
		// spew.Dump(resp)

		// Maquina 1 - Fluxo out(M1Fo)
		pathSourceM1Fo := "../sinis"
		pathDestM1Fo := "../trans"
		filesM1Fo := TaskPathfinderOfFilesModifieds(pathSourceM1Fo)
		_ = TaskCopyFiles(filesM1Fo, pathSourceM1Fo, pathDestM1Fo)
		resp := TaskCommitFilesModified(pathDestM1Fo)
		spew.Dump(resp)

		// Maquina 1 - Fluxo in(M1Fi)
		pathSourceM1Fi := "../trans"
		pathDestM1Fi := "../sinis"
		cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
		spew.Dump(cmdStrPullM1Fi)
		files := TaskPathfinderOfFilesModifieds(pathSourceM1Fi)
		_ = TaskCopyFiles(files, pathSourceM1Fi, pathDestM1Fi)
		spew.Dump(files)

		// Maquina 2 - Fluxo in(M2Fi)
		pathSourceM2Fi := "../trans"
		cmdStrPullM2Fi := gitcommand.GitPull(pathSourceM2Fi)
		spew.Dump(cmdStrPullM2Fi)

		// Maquina 2 - Fluxo out(M2Fo)
		pathSourceM2Fo := "../trans"
		respM2Fo := TaskCommitFilesModified(pathSourceM2Fo)
		spew.Dump(respM2Fo)

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
