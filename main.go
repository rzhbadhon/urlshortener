package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Server running")
	},

)

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}


}
