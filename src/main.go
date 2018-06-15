package main

import (
	"net/http"

	"./chat"
)

func main() {

	// websocket server
	server := chat.NewServer()
	go server.Listen()
	http.HandleFunc("/messages", handleHomePage)
	http.HandleFunc("/", handleHomePage)

	http.ListenAndServe("localhost:9999", nil)

}

func handleHomePage(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "chat.html")
}
