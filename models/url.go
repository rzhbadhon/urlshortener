package models

import "time"

type UrlRequest struct {
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type ShortUrl struct{
	ShortUrl string `json:"short_url"`
	ExpireAt string `json:"expire_at"`
}
