package main

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	http.Handle("/", web.NewRouter())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
