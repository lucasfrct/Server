package reverse

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	log.Println(url.Scheme)

	proxy.ServeHTTP(res, req)
}
