package command

import (
	"errors"
	"fmt"
	"os"
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
	cmdarrAux := []string{}
	pathTemp := ""
	for i := range cmdarr {
		if !strings.Contains(cmdarr[i], string(os.PathSeparator)) {
			cmdarrAux = append(cmdarrAux, cmdarr[i])
			continue
		}
		pathTemp = fmt.Sprintf("%s %s", pathTemp, cmdarr[i])
		path := strings.TrimSpace(pathTemp)
		if _, err := os.Stat(path); err == nil {
			cmdarrAux = append(cmdarrAux, path)
			pathTemp = ""
		}
	}

	cmd := exec.Command(cmdarrAux[0], cmdarrAux[1:]...)

	cmd.Dir = pathAbs
	cmd.Env = append(cmd.Environ(), "POSIXLY_CORRECT=1")

	cmdByte, err := cmd.Output()
	if err != nil {
		panic(command)
	}

	cmdStr = string(cmdByte)

	return
}

func Copy(pathSouce, pathDest string) error {

	if pathSouce == "" {
		return errors.New("path de origem vazio")
	}

	if pathDest == "" {
		return errors.New("path de destino vazio")
	}

	return cp.Copy(pathSouce, pathDest)
}
