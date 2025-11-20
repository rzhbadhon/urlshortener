package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rzhbadhon/urlshortener/models"
)

func (h *Handler) GetText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post request alowed", http.StatusUnauthorized)
		return
	}

	var code models.ShortUrl
	// decode the txt
	json.NewDecoder(r.Body).Decode(&code)
	// get the url from db
	query := "SELECT text FROM urls WHERE id = $1"

	err := h.Db.QueryRow(query, code.Id).Scan(&code.Texts)
	if err != nil {
		fmt.Println("DB insert error:", err)
		http.Error(w, "Failed to insert url into db", http.StatusInternalServerError)
		return
	}
	// send response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(code.Texts)

}
