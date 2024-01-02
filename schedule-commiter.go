package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
	"github.com/lucasfrct/servertools/tasks"
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

func init() {

	content, err := Read("./schedule-commiter.json")
	if err != nil {
		panic("O Arquivo ./schedule-commiter.json não está presente na pasta atual")
	}

	var config *Config = new(Config)
	err = json.Unmarshal([]byte(content), config)
	if err != nil {
		panic("O Arquivo ./schedule-commiter.json está com json mal formatado")
	}

	fliesModified := tasks.PathfinderOfFilesModifieds(config.Source)
	filesCopied := tasks.CopyFiles(fliesModified, config.Source, config.Destination)
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
