package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	

		
}