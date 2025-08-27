package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/", homeHandler)

	log.Fatal(server.Serve(listener))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
