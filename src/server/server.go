package server

import (
	"net/http"

	router "lucasfrct.com/router"
)

type Server struct {
	port        string
	directories []string
	directory   string
}

func (s *Server) New() {

	router.HandleRoutes()

	for _, directory := range s.directories {
		http.Handle("/"+directory, http.StripPrefix("/"+directory, http.FileServer(http.Dir(directory))))
	}

	http.ListenAndServe(":8080", nil)
}

func (s *Server) AddDirectry(directory string) {
	s.directories = append(s.directories, directory)
}
