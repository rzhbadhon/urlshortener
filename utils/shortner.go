package utils

import "github.com/rzhbadhon/urlshortener/models"

func Shortner(reqUrl string) (string, error) {
	var shortUrl models.ShortUrl
	shortUrl.ShortUrl = reqUrl

	return shortUrl.ShortUrl, nil
}