package engine

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2/log"

	"github.com/meanii/imgax/config"
	"github.com/meanii/imgax/utils"
)

type RembgEngine struct {
	RmbgUrl            string `query:"rembg"`
	RmbgApiUrl         url.URL
	ProcessedFGImage   string
	BackgroundImageUrl string `query:"bgimg"`
	BackgroundColor    string `query:"bgcol"`
}

func (r *RembgEngine) Rmbg() (string, error) {
	rembgUrl, err := url.Parse(config.Env.RembgUrl)
	if err != nil {
		log.Errorf("Error parsing rembg url: %s", err)
		return "", err
	}

	rembgUrl.Path = "/api/remove"
	q := rembgUrl.Query()
	q.Set("url", r.RmbgUrl)
	rembgUrl.RawQuery = q.Encode()
	log.Infof("Rembg Engine url: %s", rembgUrl.String())

	r.RmbgApiUrl = *rembgUrl

	rembgUrlFilename := strings.Split(filepath.Base(r.RmbgUrl), ".")[0] + ".png"

	savedPath, err := utils.Downloader(
		r.RmbgApiUrl.String(),
		rembgUrlFilename,
	)
	if err != nil {
		log.Errorf("Error downloading file: %s", err)
		return "", err
	}
	return savedPath, nil
}

func (r *RembgEngine) SetBackgroundColor() (string, error) {
	bgcfilename := fmt.Sprintf("%s.png", r.BackgroundColor)
	log.Infof("Background color: %s", r.BackgroundColor)

	imageEngine := ImageEngine{}
	imageEngine.File = filepath.Join(
		filepath.Dir(r.ProcessedFGImage),
		bgcfilename,
	)
	imageEngine.BackgroundColor = imageEngine.HexColor(r.BackgroundColor)
	_, err := imageEngine.SetForegroundImage(r.ProcessedFGImage)
	if err != nil {
		return "", err
	}
	imageInfo, _ := imageEngine.DecodeImage(r.ProcessedFGImage)
	imageEngine.Height = imageInfo.Bounds().Dy()
	imageEngine.Width = imageInfo.Bounds().Dx()
	imageEngine.CreateImage()
	processedImage, err := imageEngine.MergeFBG()
	if err != nil {
		return "", err
	}
	return processedImage, nil
}

func (r *RembgEngine) Do() (string, error) {
	log.Infof("Called RembgEngine: %s", r.RmbgUrl)

	savedPath, err := r.Rmbg()
	if err != nil {
		return "", err
	}
	r.ProcessedFGImage = savedPath

	if r.BackgroundColor != "" {
		savedPath, err = r.SetBackgroundColor()
		if err != nil {
			return "", err
		}
		return savedPath, nil
	}

	return savedPath, nil
}
