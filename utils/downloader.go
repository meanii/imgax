package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func Downloader(url string, filename string) (string, error) {
	log.Infof("Downloading image request: %s", url)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	savedPath, err := SaveFile(filename, res.Body)
	if err != nil {
		return "", err
	}
	return savedPath, nil
}
