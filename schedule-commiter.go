package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/command"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

type Config struct {
	Destination     string `json:"destination"`
	Source          string `json:"source"`
	SecondsInterval int    `json:"secondsInterval"` // em segundos
}

func Read(path string) (content string, err error) {
	content = ""

	readedFile, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return
	}

	content = string(readedFile)
	return
}

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
}

func main() {

	content, err := Read("./schedule-commiter.json")
	if err != nil {
		panic("O Arquivo ./schedule-commiter.json não está presente na pasta atual")
	}

	var config *Config = new(Config)
	err = json.Unmarshal([]byte(content), config)
	if err != nil {
		panic("O Arquivo ./schedule-commiter.json está com json mal formatado")
	}

	fliesModified := TaskPathfinderOfFilesModifieds(config.Source)
	filesCopied := TaskCopyFiles(fliesModified, config.Source, config.Destination)
	spew.Dump(filesCopied)

	// done := make(chan bool)
	// var scheduleTime time.Duration = time.Duration(config.Interval)

	// done = schedule.Schedule(func() {

	// 	filesM1Fo := TaskPathfinderOfFilesModifieds(pathSourceM1Fo)
	// 	filesM1FoCopied := TaskCopyFiles(filesM1Fo, pathSourceM1Fo, pathDestM1Fo)
	// 	resp := TaskCommitFilesModified(pathDestM1Fo)
	// 	spew.Dump(filesM1FoCopied, resp)

	// 	// Maquina 1 - Fluxo in(M1Fi)
	// 	pathSourceM1Fi := "../transfering"
	// 	pathDestM1Fi := "../sinis"
	// 	// cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
	// 	filesM1Fi := TaskPathfinderOfFilesModifieds(pathSourceM1Fi)
	// 	filesM1FiCopied := TaskCopyFiles(filesM1Fi, pathSourceM1Fi, pathDestM1Fi)
	// 	spew.Dump(filesM1FiCopied)

	// Maquina 2 - Fluxo in(M2Fi)
	pathSourceM2Fi := "../transfer"
	cmdStrPullM2Fi := gitcommand.GitPull(pathSourceM2Fi)
	spew.Dump(cmdStrPullM2Fi)

	// 	// Maquina 2 - Fluxo out(M2Fo)
	// 	pathSourceM2Fo := "../transfering"
	// 	respM2Fo := TaskCommitFilesModified(pathSourceM2Fo)
	// 	spew.Dump(respM2Fo)

	// }, scheduleTime*time.Second)

	// go func() {
	// 	time.Sleep(1 * time.Minute)
	// 	done <- true
	// }()

	// select {
	// case v1 := <-done:
	// 	fmt.Println("END: ", v1)
	// }

}
