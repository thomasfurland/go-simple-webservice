package handlers

import "net/http"

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
