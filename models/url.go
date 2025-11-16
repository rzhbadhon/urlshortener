package models

import "time"

type UrlRequest struct {
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	Texts string `json:"text"`
}

type ShortUrl struct{
	ShortUrl string `json:"short_url"`
	ExpireAt string `json:"expire_at"`
	Texts string `json:"text"`
}
