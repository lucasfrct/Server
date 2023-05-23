package router

import (
	"net/http"
)

func Assets() {
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
}

func Directories(directories []string) {
	for _, directory := range directories {
		http.Handle("/"+directory, http.StripPrefix("/"+directory, http.FileServer(http.Dir(directory))))
	}
}
