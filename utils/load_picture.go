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

type ImgCar struct {
	sprite *pixel.Sprite
	Id int
	enterating bool
	pos pixel.Vec
}

func NewImgCar(sprite *pixel.Sprite, Id int, state bool, pos pixel.Vec) *ImgCar {
	return &ImgCar{
		sprite: sprite,
		Id: Id,
		enterating: state,
		pos: pos,
	}
}

func (ic *ImgCar) SetData(sprite *pixel.Sprite, Id int, state bool, pos pixel.Vec) {
	ic.sprite = sprite
	ic.Id = Id
	ic.enterating = state
	ic.pos = pos
}

func (ic *ImgCar) GetSprite() *pixel.Sprite {
	return ic.sprite
}

func (ic *ImgCar) GetPos() pixel.Vec {
	return ic.pos
}

func (ic *ImgCar) GetId() int {
	return ic.Id
}

func (ic *ImgCar) GetStatus() bool {
	return ic.enterating
}

func (ic *ImgCar) GetData() *ImgCar {
	return ic
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