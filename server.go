package main

import (
	"net/http"

	router "localhost.com/router"
)

func Init() {

	router.HandleRoutes()
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}
