package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/my-wallet", func(writer http.ResponseWriter, request *http.Request) {

		//Membuat object wallet
		exampleWallet := Wallet{
			Name:    "Ari Susanto",
			Balance: 1200,
		}

		//encode ke json
		byteWallet, err := json.Marshal(exampleWallet)
		if err != nil {
			panic(err)
		}

		//kirim ke client sebagai json
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(byteWallet)
		if err != nil {
			panic(err)
		}

	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: muxHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type Wallet struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
