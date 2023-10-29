package utils

import (
	"image"
	"os"
	"github.com/faiface/pixel"
)

type Utils struct {
	pic *pixel.Picture
	sprite *pixel.Sprite

}

func (u *Utils) LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (u *Utils) NewSprite(pic pixel.Picture, form pixel.Rect) *pixel.Sprite {
	return pixel.NewSprite(pic, form)
}