package handlers

import (
	"net/http"
	"time"
)

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request){
	shortCode := r.URL.Path[len("/"):]

	var originalUrl string
	var expireAt time.Time

	err := h.Db.QueryRow("SELECT original_url, expire_at FROM urls WHERE short_code=$1", shortCode).Scan(&originalUrl, &expireAt)
	if err != nil{
		http.Error(w, "Short URL not found", http.StatusNotFound)
        return
	}

	if time.Now().After(expireAt){
		http.Error(w, "Short URL expired", http.StatusGone)
        return
	}

	http.Redirect(w, r, originalUrl, http.StatusFound)

}