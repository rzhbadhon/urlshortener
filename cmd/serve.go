package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rzhbadhon/urlshortener/rest/handlers"
	"github.com/rzhbadhon/urlshortener/rest/middlewares"
)

func Server() {
	//"user=postgres password=1212 dbname=authentication sslmode=disable"
	connStr := "user=postgres password=1212 dbname=urlshortner sslmode=disable"
	if connStr == "" {
		log.Fatal("DB_URL isnt set")
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	h := handlers.NewHandler(db)

	log.Println("Database connected succssfully")

	mux := http.NewServeMux()

	mux.HandleFunc("/", http.HandlerFunc(h.RedirectURL))
	mux.HandleFunc("/shorten", http.HandlerFunc(h.ShortUrl))
	mux.HandleFunc("/text", http.HandlerFunc(h.GetText))

	handlerMux := middlewares.CorsHandler(mux)

	err = http.ListenAndServe(":5000", handlerMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
