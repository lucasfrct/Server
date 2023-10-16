package tasks

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func Push(pathDest string) {

	// time.Sleep(5 * time.Second)

	fmt.Print("Pushing to Github repository!")

	resp := CommitFilesModified(pathDest)
	spew.Dump(resp)

}
