package main

import "net/http"

func main() {

	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/my-wallet", SetCookiesWeb)
	muxHandler.HandleFunc("/my-cookie", GetCookiesWeb)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: muxHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func SetCookiesWeb(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:  "My-Cookie",
		Value: request.Host,
		Path:  "/my-wallet",
	}

	http.SetCookie(writer, &cookie)
}

func GetCookiesWeb(writer http.ResponseWriter, request *http.Request) {
	resultCookie, err := request.Cookie("My-Cookie")
	if err != nil {
		writer.Write([]byte("No Cookie Available"))
	} else {
		result := resultCookie.Value
		writer.Write([]byte(result))
	}
}
