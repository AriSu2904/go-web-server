package main

import (
	"fmt"
	"net/http"
)

func main() {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World!")
	}

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
