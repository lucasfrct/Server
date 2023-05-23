package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ServeReverseProxy(targetpoint string, res http.ResponseWriter, req *http.Request) {

	url, _ := url.Parse(targetpoint)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	log.Println(url.Scheme)

	proxy.ServeHTTP(res, req)
}

func HandleRequestAndRedirect(targetpoint string) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		ServeReverseProxy(targetpoint, res, req)
	}
}

func New(entrypoint, targetpoint string) {

	log.Println("=======================================================")
	log.Print("Start redirection: ENTRYPOINT(:8080)->TARGETPOINT(:4200)")
	log.Println("=======================================================")
	log.Println("")

	http.HandleFunc("/", HandleRequestAndRedirect(targetpoint))
	_ = http.ListenAndServe(entrypoint, nil)
}
