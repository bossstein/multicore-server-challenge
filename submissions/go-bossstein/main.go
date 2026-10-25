package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
	})

	http.ListenAndServe("127.0.0.1:3000", mux)
}
