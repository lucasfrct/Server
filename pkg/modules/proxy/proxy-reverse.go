package proxy

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type Proxy struct {
	EntrypointUrl  string
	TargetpointUrl string
}

func (p *Proxy) Redirect(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(p.TargetpointUrl)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Manipulação personalizada da solicitação antes de encaminhá-la para o servidor SSR
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Lê o corpo da resposta do servidor SSR
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Realiza qualquer manipulação necessária no conteúdo do corpo da resposta
		// Aqui você pode fazer alterações específicas no HTML retornado pelo servidor SSR

		// Atualiza o conteúdo do corpo da resposta
		newBody := []byte("Manipulated SSR content: " + string(body))
		resp.Body = ioutil.NopCloser(strings.NewReader(string(newBody)))
		resp.ContentLength = int64(len(newBody))

		return nil
	}

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("Access-Control-Allow-Origin", "*")
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
