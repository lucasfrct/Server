package staticfile

import "net/http"

func StaticFile(path string, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path)
}
