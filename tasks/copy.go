package tasks

import (
	"github.com/lucasfrct/servertools/pkg/modules/command"
)

func Copy(pathSource, pathDest string) (err error) {
	return command.Copy(pathSource, pathDest)
}
