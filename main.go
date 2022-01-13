package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			writer.WriteHeader(http.StatusOK)
			fmt.Fprint(writer, "pong")
		default:
			http.NotFound(writer, request)
		}
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}

