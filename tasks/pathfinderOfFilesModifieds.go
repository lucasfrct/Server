package tasks

import (
	"path/filepath"

	"github.com/lucasfrct/servertools/pkg/modules/command"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

func PathfinderOfFilesModifieds(pathSource string) []string {

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
