package server

import (
	"net/http"

	"github.com/lucasfrct/servertools/pkg/routes"
)

type Server struct {
	port        string
	directories []string
}

func (s *Server) New() {

	s.port = ":8080"

	routes.Assets()
	routes.Directories(s.directories)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) AddDirectory(directory string) {
	s.directories = append(s.directories, directory)
}
