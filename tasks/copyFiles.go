package tasks

import (
	"os"
	"path/filepath"
	"strings"
)

func CopyFiles(files []string, pathSource, pathDest string) []string {

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
		Copy(source, dest)
		files[i] = dest
	}

	return files
}
