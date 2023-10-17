package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello from endpoint hello")
	})

	mux.HandleFunc("/greeting", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello from greeting endpoint")
	})

	mux.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, request.Method)
		fmt.Fprintf(writer, request.RequestURI)
	})

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
