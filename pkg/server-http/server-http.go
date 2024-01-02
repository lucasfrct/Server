package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	Http              *http.ServeMux
	Port              int
	Host              string
	Address           string
	Protocol          string
	Paths             []string
	StaticDirectories []string
}

func (s *Server) New() *Server {
	s.Environment()
	s.NewHttp()
	return s
}

func (s *Server) Environment() *Server {
	s.Port = 8080
	s.Host = "0.0.0.0"
	s.Protocol = "http"
	s.Address = fmt.Sprintf("%v:%v", s.Host, s.Port)
	return s
}

func (s *Server) NewHttp() *Server {
	s.Http = http.NewServeMux()
	return s
}

func (s *Server) AddPort(port int) *Server {
	s.Port = port
	s.Address = fmt.Sprintf("%v:%v", s.Host, s.Port)
	return s
}

func (s *Server) AddHost(host string) *Server {
	s.Host = host
	s.Address = fmt.Sprintf("%v:%v", s.Host, s.Port)
	return s
}

func (s *Server) AddDirectory(directory string) (*Server, error) {

	directory = strings.TrimSpace(directory)

	directory, err := filepath.Abs(directory)
	if err != nil {
		log.Printf("Erro. diretório inválido. %v", err.Error())
		return nil, err
	}

	s.StaticDirectories = append(s.StaticDirectories, directory)
	return s, nil
}

func (s *Server) RunStaticDirectories() *Server {

	for i := range s.StaticDirectories {

		dir := s.StaticDirectories[i]
		uri := fmt.Sprintf("/%s/", filepath.Base(dir))

		s.Http.Handle(uri, http.StripPrefix(uri, http.FileServer(http.Dir(dir))))
		s.Paths = append(s.Paths, uri)
		time.Sleep(100 * time.Millisecond)
	}

	return s
}

func (s *Server) Listen() error {
	err := http.ListenAndServe(s.Address, cors.Default().Handler(s.Http))
	if err != nil {
		log.Printf("Erro. Não foi possivel iniciar o servidor. %v", err.Error())
		return err
	}

	return nil
}

func NewServerHTTPStatic(directories []string) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite a porta do servidor que quer levantar: ex.: (8080, 81, 90)")
	scanner.Scan()
	port := scanner.Text()
	fmt.Println("")

	server := new(Server).New()
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("A porta passada não é um número válido.")
		return
	}

	server.AddPort(portInt)

	for i := range directories {
		server.AddDirectory(directories[i])
	}

	server.RunStaticDirectories()

	fmt.Println("=======================================  SERVER HTTP STATIC ========================================================")
	fmt.Println("====================================================================================================================")
	fmt.Println("")
	fmt.Printf("         Iniciando Servidor: ENTRYPOINT(%v)", server.Address)
	fmt.Println("")
	fmt.Println("")
	for i := range server.Paths {
		msg := fmt.Sprintf("         *  -->  Wathcher Directory: %v%s -> %s", server.Address, server.Paths[i], server.StaticDirectories[i])
		fmt.Println(msg)
	}
	fmt.Println("")
	fmt.Println("====================================================================================================================")
	fmt.Println("====================================================================================================================")

	err = server.Listen()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	NewServerHTTPStatic([]string{
		`C:\Users\lucas\Documents\github.com\lucasfrct\servertools\assets`,
		`C:\Users\lucas\Documents\github.com\lucasfrct\servertools\tasks`,
	})
}
