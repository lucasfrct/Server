package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/lucasfrct/servertools/pkg/modules/schedule"
)

func main() {
	// proxy.Reverse()
	// 	schedule.Schedule()
	done := schedule.Schedule(func() {
		println("Schedule: \n")

		command := ""
		cmdarr := strings.Split(command, " ")
		cmd, err := exec.Command(cmdarr[0], cmdarr[1:]...).Output()
		if err != nil {
			fmt.Printf("Erro: %v \n", err)
		}

		fmt.Printf("%s \n", string(cmd))

	}, 5*time.Second)

	select {
	case v1 := <-done:
		fmt.Println("got: ", v1)
	}

	done <- false

}
