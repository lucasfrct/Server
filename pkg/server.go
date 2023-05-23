package server

import (
	"net/http"

	"github.com/lucasfrct/servertools/pkg/router"
)

type Server struct {
	port        string
	directories []string
	directory   string
}

func (s *Server) New() {

	router.Assets()
	router.Directories(s.directories)
	http.ListenAndServe(":8080", nil)
}

func (s *Server) AddDirectory(directory string) {
	s.directories = append(s.directories, directory)
}
