package main

import (
	"os"
)

func RootDir() string {
	dir, _ := os.Getwd()
	return dir
}
