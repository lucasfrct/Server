package router

import (
	"net/http"

	account "localhost.com/account"
	login "localhost.com/login"
)

func HandleRoutes() {

	http.HandleFunc("/", login.Index)
	http.HandleFunc("/account", account.Index)
}

func Init() {

	HandleRoutes()
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}
