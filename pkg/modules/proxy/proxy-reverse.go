package proxy

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Proxy struct {
	EntrypointUrl  string
	TargetpointUrl string
}

func (p *Proxy) Redirect(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(p.TargetpointUrl)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func (p *Proxy) Listen() error {
	http.HandleFunc("/", p.Redirect)
	return http.ListenAndServe(p.EntrypointUrl, nil)
}

func Reverse() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite a porta do servidor que quer levantar: ex.: (8080, 81, 90)")
	scanner.Scan()
	portEntrypointUrl := scanner.Text()
	entrypointUrl := fmt.Sprintf("0.0.0.0:%v", portEntrypointUrl)
	fmt.Println("")

	fmt.Println("Digite a Url que quer copiar/redirecionar:")
	scanner.Scan()
	targetpointUrl := scanner.Text()
	fmt.Println("")

	fmt.Println("============================================  PROXY REVERSO ========================================================")
	fmt.Println("====================================================================================================================")
	fmt.Println("")
	fmt.Printf("         Iniciando redirecionamento: ENTRYPOINT(%v)->TARGETPOINT(%v)", entrypointUrl, targetpointUrl)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("====================================================================================================================")
	fmt.Println("====================================================================================================================")

	prx := Proxy{EntrypointUrl: entrypointUrl, TargetpointUrl: targetpointUrl}
	_ = prx.Listen()
}
