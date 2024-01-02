package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/rs/cors"
)

type Server struct {
	Http     *http.ServeMux
	Port     int
	Host     string
	Address  string
	Protocol string
	Target   string
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

func (s *Server) Redirect(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(s.Target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	for key, values := range req.Header {
		for _, value := range values {
			proxy.Transport = &http.Transport{}
			proxy.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = true
			res.Header().Add(key, value)
		}
	}

	// Verifica se o cliente suporta streaming de texto
	if req.Header.Get("Accept") == "text/stream" {
		req.Header.Set("Content-Type", "text/stream")
		req.Header.Set("Transfer-Encoding", "chunked")
	}

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func (s *Server) Listen() error {
	http.HandleFunc("/", s.Redirect)
	err := http.ListenAndServe(s.Address, cors.Default().Handler(s.Http))
	if err != nil {
		log.Printf("Erro. Não foi possivel iniciar o servidor proxy. %v", err.Error())
		return err
	}

	return nil
}

func NewServerOfProxyReverse() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite a porta do servidor que quer levantar: ex.: (8080, 81, 90)")
	scanner.Scan()
	port := scanner.Text()
	fmt.Println("")

	fmt.Println("Digite a Url do servidor que vai receber as requisições redirecionadas:")
	scanner.Scan()
	targetpointUrl := scanner.Text()
	fmt.Println("")

	server := new(Server).New()
	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("A porta passada não é um número válido.")
		return
	}

	server.AddPort(portInt)
	server.Target = strings.TrimSpace(targetpointUrl)

	fmt.Println("============================================  PROXY REVERSO ========================================================")
	fmt.Println("====================================================================================================================")
	fmt.Println("")
	fmt.Printf("         Iniciando redirecionamento: ENTRYPOINT(%v)->TARGETPOINT(%v)", server.Address, server.Target)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("====================================================================================================================")
	fmt.Println("====================================================================================================================")

	err = server.Listen()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	NewServerOfProxyReverse()
}
