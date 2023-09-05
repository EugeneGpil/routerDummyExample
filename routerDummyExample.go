package main

import (
	"net/http"

	"github.com/EugeneGpil/router"
)

func main() {
	method := http.MethodGet
	url := "hello/world"
	name := "hello.world"

	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!\n"))
	}

	router.AddRoute(method, url, handler, name)

	mux := http.NewServeMux()

	router.DefineRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
