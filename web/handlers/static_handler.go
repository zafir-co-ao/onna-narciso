package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func NewStaticHandler() func(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Executable()
	cwd = path.Dir(cwd)

	indexFile := getPath(cwd, "static/index.html")

	staticDir := getPath(cwd, "static")
	staticFS := http.Dir(staticDir)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, indexFile)
			return
		}

		http.StripPrefix("/", http.FileServer(staticFS)).ServeHTTP(w, r)

	}
}

func getPath(cwd string, p string) string {
	return fmt.Sprintf("%s/%s", cwd, p)
}
