package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/command"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
	"github.com/lucasfrct/servertools/pkg/modules/schedule"
)

func TaskPathfinderOfFilesModifieds(pathSource string) []string {

	pathAbs, err := filepath.Abs(pathSource)
	if err != nil {
		panic(err)
	}

	gitFiles, err := command.Exec(pathAbs, gitcommand.GitListFiles())
	if err != nil {
		panic(err)
	}

	files := gitcommand.GitListFilesModified(gitFiles)

	for i := range files {
		files[i] = filepath.Join(pathAbs, files[i])
	}

	return files
}

func TaskCopy(pathSource, pathDest string) error {
	return command.Copy(pathSource, pathDest)
}

func TaskCopyFiles(files []string, pathSource, pathDest string) []string {

	pathSourceAbs, err := filepath.Abs(pathSource)
	if err != nil {
		panic(err)
	}

	pathDestAbs, err := filepath.Abs(pathDest)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(pathDestAbs, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for i := range files {
		source := files[i]
		dest := filepath.Join(pathDestAbs, strings.Split(files[i], pathSourceAbs)[1])
		TaskCopy(source, dest)
		files[i] = dest
	}

	return files
}

func TaskCommitFilesModified(pathSource string) string {
	files := TaskPathfinderOfFilesModifieds(pathSource)
	cmd := gitcommand.GitProcessToCommit(files)

	var resp string = ""
	commands := strings.Split(cmd, ";")
	for i := range commands {
		c := strings.TrimSpace(commands[i])

		if len(c) == 0 || c == "" {
			continue
		}

		fmt.Println(" -- > RUN: ", c)
		r, err := command.Exec(pathSource, c)
		if err != nil {
			continue
		}

		resp += r

	}

	return resp
}

func TaskSyncProject(source, dest string) {
	err := TaskCopy(source, dest)
	if err != nil {
		panic("Errro ao tentar copiar arquivos")
	}

	fmt.Println(fmt.Sprintf("* Arquivos copiados: (%s) -> (%s)", source, dest))
	resp := TaskCommitFilesModified(dest)
	fmt.Println(fmt.Sprintf("* Projeto Commitado (%s): %s", dest, resp))
	resp = gitcommand.GitPull(dest)
	fmt.Println(fmt.Sprintf("* Projeto sincronizado (%s): %s", dest, resp))

	// tira a diferenca entre dois diretórios
}

func main() {

	done := make(chan bool)
	var scheduleTime time.Duration = 10

	done = schedule.Schedule(func() {

		// Maquina 1 - Fluxo out(M1Fo)
		pathSourceM1Fo := "../sinis"
		pathDestM1Fo := "../transfering"
		filesM1Fo := TaskPathfinderOfFilesModifieds(pathSourceM1Fo)
		filesM1FoCopied := TaskCopyFiles(filesM1Fo, pathSourceM1Fo, pathDestM1Fo)
		// resp := TaskCommitFilesModified(pathDestM1Fo)
		spew.Dump(filesM1FoCopied)

		// Maquina 1 - Fluxo in(M1Fi)
		pathSourceM1Fi := "../transfering"
		pathDestM1Fi := "../sinis"
		// cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
		filesM1Fi := TaskPathfinderOfFilesModifieds(pathSourceM1Fi)
		filesM1FiCopied := TaskCopyFiles(filesM1Fi, pathSourceM1Fi, pathDestM1Fi)
		spew.Dump(filesM1FiCopied)

		// // Maquina 2 - Fluxo in(M2Fi)
		// pathSourceM2Fi := "../transfering"
		// cmdStrPullM2Fi := gitcommand.GitPull(pathSourceM2Fi)
		// spew.Dump(cmdStrPullM2Fi)

		// // Maquina 2 - Fluxo out(M2Fo)
		// pathSourceM2Fo := "../transfering"
		// respM2Fo := TaskCommitFilesModified(pathSourceM2Fo)
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
