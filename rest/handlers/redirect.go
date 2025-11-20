package handlers

import (
	"net/http"
	"strings"
	"time"
)

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request){
	shortCode := r.URL.Path[len("/"):]

	
	var originalUrl string
	var expireAt time.Time // var to set the expire time

	err := h.Db.QueryRow("SELECT original_url, expire_at FROM urls WHERE short_code=$1", shortCode).Scan(&originalUrl, &expireAt)
	if err != nil{
		http.Error(w, "Short URL not found", http.StatusNotFound)
        return
	}

	// check if the url is expired or not
	if time.Now().After(expireAt){
		http.Error(w, "Short URL expired", http.StatusGone)
        return
	}

	if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://"){
		originalUrl = "https://" + originalUrl
	}
	// to redirect the user to the destination url
	http.Redirect(w, r, originalUrl, http.StatusFound)

}