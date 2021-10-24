package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	log.Println(url.Scheme)

	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := "http://localhost:4200"
	serveReverseProxy(url, res, req)
}

func main() {
	log.Println("=======================================================")
	log.Print("start redirect :8080->:4200")
	log.Println("=======================================================")
	log.Println("")

	http.HandleFunc("/", handleRequestAndRedirect)
	http.ListenAndServe(":8080", nil)
}
