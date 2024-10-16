package routes

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

var cwd string

func init() {
	cwd, _ = os.Executable()
	cwd = path.Dir(cwd)
}

func New() *http.ServeMux {

	assetsDir := getPath("assets")
	staticDir := getPath("static")
	indexFile := getPath("static/index.html")

	assetsFS := http.Dir(assetsDir)
	staticFS := http.Dir(staticDir)

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(assetsFS)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, indexFile)
			return
		}

		http.StripPrefix("/", http.FileServer(staticFS)).ServeHTTP(w, r)

	})

	return mux
}

func getPath(p string) string {
	return fmt.Sprintf("%s/%s", cwd, p)
}
