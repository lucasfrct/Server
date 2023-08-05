package command

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

func Execute(command string) (cmd string, err error) {

	cmdarr := strings.Split(command, " ")
	cmdByte, err := exec.Command(cmdarr[0], cmdarr[1:]...).Output()
	if err != nil {
		return
	}

	cmd = string(cmdByte)

	return
}

func Exec(path string, command string) (cmdStr string, err error) {

	pathAbs, err := filepath.Abs(path)
	if err != nil {
		panic(path)
	}

	cmdarr := strings.Split(command, " ")
	cmd := exec.Command(cmdarr[0], cmdarr[1:]...)

	cmd.Dir = pathAbs
	cmd.Env = append(cmd.Environ(), "POSIXLY_CORRECT=1")

	cmdByte, err := cmd.Output()
	if err != nil {
		panic(command)
	}

	cmdStr = string(cmdByte)
	return
}

func Copy(pathSouce, pathDest string) (err error) {

	if pathSouce == "" {
		return errors.New("path de origem vazio")
	}

	if pathDest == "" {
		return errors.New("path de destino vazio")
	}

	err = cp.Copy(pathSouce, pathDest)
	return
}
