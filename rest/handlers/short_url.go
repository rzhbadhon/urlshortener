package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rzhbadhon/urlshortener/models"
	"github.com/rzhbadhon/urlshortener/utils"
)

type Handler struct {
	Db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		Db: db,
	}
}

func (h *Handler) ShortUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqUrl models.UrlRequest
	err := json.NewDecoder(r.Body).Decode(&reqUrl)
	if err != nil {
		fmt.Println("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var id int
	expireAt := time.Now().Add(48 * time.Hour)
	err = h.Db.QueryRow("INSERT INTO urls (original_url, expire_at, text) VALUES ($1, $2, $3) RETURNING id", reqUrl.URL, expireAt, reqUrl.Texts).Scan(&id)
	if err != nil {
		fmt.Println("DB insert error:", err)
		http.Error(w, "Failed to insert url into db", http.StatusInternalServerError)
		return
	}

	shortCode := utils.Shortner(id)
	_, err = h.Db.Exec("UPDATE urls SET short_code=$1 WHERE id=$2", shortCode, id)
	if err != nil {
		http.Error(w, "Failed to update short code", http.StatusInternalServerError)
		return
	}

	var sendUrl models.ShortUrl
	sendUrl.ExpireAt = expireAt.Format("2006-01-02 15:04:05")

	// we are sending the full shorturl to the user... not just the shortcode..
	sendUrl.ShortUrl = "http://localhost:5000/" + shortCode
	sendUrl.Texts = reqUrl.Texts
	sendUrl.Code = shortCode
	sendUrl.Id = id

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sendUrl)

}
