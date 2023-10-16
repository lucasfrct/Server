package tasks

import "github.com/lucasfrct/servertools/pkg/modules/command"

func Copy(pathSource, pathDest string) error {
	return command.Copy(pathSource, pathDest)
}
