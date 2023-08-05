package routes

import (
	"fmt"
	"net/http"
)

func Assets() {
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
}

func Directories(directories []string) {
	for i := range directories {
		dir := directories[i]
		routeDir := fmt.Sprintf("/%v", dir)
		http.Handle(routeDir, http.StripPrefix(routeDir, http.FileServer(http.Dir(dir))))
	}
}
