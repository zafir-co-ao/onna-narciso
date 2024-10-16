package main

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/routes"
)

func main() {
	http.Handle("/", routes.New())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
