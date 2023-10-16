package tasks

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

func Pull(pathSourceM1Fi string) {

	fmt.Print("Pulling from Github repository! \n")

	cmdStrPullM1Fi := gitcommand.GitPull(pathSourceM1Fi)
	spew.Dump(cmdStrPullM1Fi)

}
