package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Proxy struct {
	urlObject string
	urlTarget string
}

func (p *Proxy) New(urlObject, urlTarget string) {
	p.urlObject = urlObject
	p.urlTarget = urlTarget

	http.HandleFunc("/", p.redirect)
	http.ListenAndServe(p.urlObject, nil)
}

func (p *Proxy) redirect(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(p.urlTarget)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// log.Println(url.Scheme)

	proxy.ServeHTTP(res, req)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite a porta do servidor que quer levantar:")
	scanner.Scan()
	portOnject := scanner.Text()
	fmt.Println("")

	fmt.Println("Digite a Url que quer copiar/redirecionar:")
	scanner.Scan()
	url := scanner.Text()
	fmt.Println("")

	fmt.Println("============================================  PROXY REVERSO ========================================================")
	fmt.Println("====================================================================================================================")
	fmt.Println("")
	fmt.Printf("    Redirecionando Servicor [ %s  --->  %s ] ", url, portOnject)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("====================================================================================================================")
	fmt.Println("====================================================================================================================")

	// prx := Proxy{urlObject: portOnject, urlTarget: url}
	// prx.New(portOnject, url)
}
