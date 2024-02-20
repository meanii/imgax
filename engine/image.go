package engine

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
)

type ImageEngine struct {
	File            string
	BackgroundColor color.RGBA
	BackgroundImage string
	ForegroundImage image.Image
	Height          int
	Width           int
}

// DecodeImage decodes the image file and returns the image.Image object
func (i *ImageEngine) DecodeImage(file string) (image.Image, error) {
	log.Infof("Decoding image: %s", file)
	imgfd, err := os.Open(file)
	if err != nil {
		log.Errorf("Error opening file: %s", err)
		return nil, err
	}

	defer imgfd.Close()
	imgExt := filepath.Ext(file)
	switch imgExt {
	case ".png":
		imgf, err := png.Decode(imgfd)
		if err != nil {
			log.Errorf("Error decoding file: %s", err)
			return nil, err
		}
		log.Infof("Image decoded: %s", file)
		return imgf, nil
	default:
		imgf, err := jpeg.Decode(imgfd)
		if err != nil {
			log.Errorf("Error decoding file: %s", err)
			return nil, err
		}
		log.Infof("Image decoded: %s", file)
		return imgf, nil
	}
}

func (i *ImageEngine) HexColor(hex string) color.RGBA {
	values, _ := strconv.ParseUint(string(hex[1:]), 16, 32)
	return color.RGBA{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
		A: 255,
	}
}

// CreateImage creates a new image with the specified width and height
func (i *ImageEngine) CreateImage() *image.RGBA {
	log.Infof("Creating new image with width: %d and height: %d", i.Width, i.Height)
	rect := image.Rect(0, 0, i.Width, i.Height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{i.BackgroundColor}, image.Point{}, draw.Src)
	return img
}

// SetForegroundImage sets the image to be used as the foreground image
func (i *ImageEngine) SetForegroundImage(img string) (image.Image, error) {
	log.Infof("Setting foreground image: %s", img)
	docodedImage, err := i.DecodeImage(img)
	if err != nil {
		return nil, err
	}
	i.ForegroundImage = docodedImage
	log.Infof("Foreground image set: %s", img)
	return docodedImage, nil
}

// MergeFBG merges the foreground and background images
func (i *ImageEngine) MergeFBG() (string, error) {
	img := i.CreateImage()
	if i.ForegroundImage == nil {
		log.Error("Foreground image not set")
		return "", errors.New("Foreground image not set")
	}
	draw.Draw(img, img.Bounds(), i.ForegroundImage, image.Point{}, draw.Over)

	file, err := os.Create(i.File)
	if err != nil {
		return "", err
	}
	defer file.Close()

	jpeg.Encode(file, img, nil)
	return i.File, nil
}
