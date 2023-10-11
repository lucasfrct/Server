package tasks

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/command"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

func CommitFilesModified(pathSource string) string {
	files := PathfinderOfFilesModifieds(pathSource)
	cmd := gitcommand.GitProcessToCommit(files)

	var resp string = ""
	commands := strings.Split(cmd, ";")
	for i := range commands {
		c := strings.TrimSpace(commands[i])
		spew.Dump("Tirando espacos: \n", c)

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
