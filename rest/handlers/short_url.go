package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rzhbadhon/urlshortener/models"
	"github.com/rzhbadhon/urlshortener/utils"
)

func ShortUrl(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqUrl models.UrlRequest
	err := json.NewDecoder(r.Body).Decode(&reqUrl)
	if err != nil{
		fmt.Println("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortUrl, err := utils.Shortner(reqUrl.URL)
	if err != nil{
		fmt.Println("Error shortning the url", http.StatusInternalServerError)
		return
	}

	var sendUrl models.ShortUrl
	sendUrl.ShortUrl = shortUrl
	expireAt := time.Now().Add(time.Hour*2)
	sendUrl.ExpireAt = expireAt.Format("2006-01-02 15:04:05")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sendUrl)
		
}