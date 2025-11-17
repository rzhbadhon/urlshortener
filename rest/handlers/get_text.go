package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rzhbadhon/urlshortener/models"
)

func(h *Handler) GetText(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only post request alowed", http.StatusUnauthorized)
		return
	}

	var code models.ShortUrl

	json.NewDecoder(r.Body).Decode(&code)

	query := "SELECT text_column FROM your_table WHERE id = 123;"
	
}