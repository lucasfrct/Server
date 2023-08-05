package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/cors"
)

type Server struct {
	Port        string
	Directories []string
	Http        *http.ServeMux
	Address     string
	Paths       []string
}

func (s *Server) NewHttp() *Server {
	s.Http = http.NewServeMux()
	return s
}

func (s *Server) AddPort(port string) *Server {
	s.Port = strings.TrimSpace(port)
	s.Address = fmt.Sprintf("0.0.0.0:%s", s.Port)
	return s
}

func (s *Server) AddDirectory(directory string) *Server {

	directory = strings.TrimSpace(directory)

	directory, err := filepath.Abs(directory)
	if err != nil {
		panic("Diretório inválido")
	}

	// directory = filepath.Clean(directory)
	s.Directories = append(s.Directories, directory)
	return s
}

func (s *Server) StaticDirectories() *Server {

	for i := range s.Directories {
		dir := s.Directories[i]
		uri := fmt.Sprintf("/%s/", filepath.Base(dir))
		s.Http.Handle(uri, http.StripPrefix(uri, http.FileServer(http.Dir(dir))))
		s.Paths = append(s.Paths, uri)
	}

	return s
}

func (s *Server) Listen() error {
	return http.ListenAndServe(s.Address, cors.Default().Handler(s.Http))
}

func NewServerHTTPStatic(directories []string) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite a porta do servidor que quer levantar: ex.: (8080, 81, 90)")
	scanner.Scan()
	port := scanner.Text()
	fmt.Println("")

	server := &Server{}
	server.AddPort(port)

	for i := range directories {
		server.AddDirectory(directories[i])
	}

	server.NewHttp().StaticDirectories()

	fmt.Println("=======================================  SERVER HTTP STATIC ========================================================")
	fmt.Println("====================================================================================================================")
	fmt.Println("")
	fmt.Printf("         Iniciando Servidor: ENTRYPOINT(%v)", server.Address)
	fmt.Println("")
	fmt.Println("")
	for i := range server.Directories {
		fmt.Println(fmt.Sprintf("         *  -->  Wathcher Directory: %s -> %s", server.Paths[i], server.Directories[i]))
	}
	fmt.Println("")
	fmt.Println("====================================================================================================================")
	fmt.Println("====================================================================================================================")

	err := server.Listen()
	if err != nil {
		panic(err)
	}
}

func main() {
	NewServerHTTPStatic([]string{
		`C:\Users\lucas\Documents\github.com\lucasfrct\servertools\assets`,
		`C:\Users\lucas\Documents\github.com\lucasfrct\servertools\pkg`,
	})
}
