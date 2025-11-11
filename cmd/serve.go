package cmd

import (
	"fmt"
	"net/http"

	"github.com/rzhbadhon/urlshortener/rest/handlers"
)

func Server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", http.HandlerFunc(handlers.ShortUrl))

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
