package tasks

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

func Pull(pathSourceM1Fi string) {

	spew.Dump("test")

	cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
	spew.Dump(cmdStrPullM1Fi)

}
