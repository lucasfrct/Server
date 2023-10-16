package tasks

import (
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/command"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

func Push(pathDest string) {

	time.Sleep(5 * time.Second)

	spew.Dump("test")

	// filesM1Fo := TaskPathfinderOfFilesModifieds(pathSource)
	// filesM1FoCopied := TaskCopyFiles(filesM1Fo, pathSource, pathDest)
	resp := TaskCommitFilesModified(pathDest)
	// spew.Dump(filesM1FoCopied)
	spew.Dump(resp)

}

// func TaskPathfinderOfFilesModifieds(pathSource string) []string {

// 	pathAbs, err := filepath.Abs(pathSource)
// 	if err != nil {
// 		panic(err)
// 	}

// 	gitFiles, err := command.Exec(pathAbs, gitcommand.GitListFiles())
// 	if err != nil {
// 		panic(err)
// 	}

// 	files := gitcommand.GitListFilesModified(gitFiles)

// 	for i := range files {
// 		files[i] = filepath.Join(pathAbs, files[i])
// 	}

// 	return files
// }

// func TaskCopyFiles(files []string, pathSource, pathDest string) []string {

// 	pathSourceAbs, err := filepath.Abs(pathSource)
// 	if err != nil {
// 		panic(err)
// 	}

// 	pathDestAbs, err := filepath.Abs(pathDest)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = os.MkdirAll(pathDestAbs, os.ModePerm)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := range files {
// 		source := files[i]
// 		dest := filepath.Join(pathDestAbs, strings.Split(files[i], pathSourceAbs)[1])
// 		TaskCopy(source, dest)
// 		files[i] = dest
// 	}

// 	return files
// }

func TaskCommitFilesModified(pathSource string) string {
	files := PathfinderOfFilesModifieds(pathSource)
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

// func TaskCopy(pathSource, pathDest string) error {
// 	return command.Copy(pathSource, pathDest)
// }
