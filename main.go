package main

import "log"

func check(e error) {
	if e != nil {
		log.Println(e)
		panic(e)
	}
}

func main() {

}
