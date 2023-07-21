package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/lucasfrct/servertools/pkg/modules/schedule"
)

func copy(command string) error {

	cmdarr := strings.Split(command, " ")
	cmd, err := exec.Command(cmdarr[0], cmdarr[1:]...).Output()
	if err != nil {
		return err
	}

	fmt.Printf("%s \n", string(cmd))
	return nil
}

func main() {
	// proxy.Reverse()
	// 	schedule.Schedule()
	done := make(chan bool)

	done = schedule.Schedule(func() {
		println("Schedule: \n")

		// Construa o comando xcopy
		command := `xcopy C:\Users\lucas\Desktop\Origem\ C:\Users\lucas\Desktop\Destino\ /s /e /y`

		err := copy(command)
		if err != nil {
			fmt.Print(err)
		}

		done <- true

	}, 5*time.Second)

	go func() {
		time.Sleep(4 * time.Second)
		done <- true
	}()

	select {
	case v1 := <-done:
		fmt.Println("got: ", v1)
	}

}
