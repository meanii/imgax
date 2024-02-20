package utils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/log"

	"github.com/meanii/imgax/config"
)

func SaveFile(filename string, data io.Reader) (string, error) {
	log.Infof("Saving file: %s", filename)
	cwd, _ := os.Getwd()
	cacheDir := filepath.Join(cwd, config.Env.CacheDir)
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		log.Warnf("Cache directory does not exist, creating: %s", cacheDir)
		os.Mkdir(cacheDir, 0755)
	}
	filePath := filepath.Join(cacheDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Errorf("Error creating file: %s", err)
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, data)
	if err != nil {
		log.Errorf("Error copying file: %s", err)
		return "", err
	}
	log.Infof("File saved: %s", filePath)
	return filePath, nil
}
