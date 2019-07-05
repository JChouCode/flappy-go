package yoshi

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"os"
)

func loadPicture(path string) (pixel.Picture, error) {
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

const gravity = 0.5

var velocity float64 = 0

var lift float64 = 30

type Yoshi struct {
	Sp *pixel.Sprite
	V  pixel.Vec
}

func New(img string) Yoshi {
	pic, err := loadPicture(img)
	if err != nil {
		panic(err)
	}
	return Yoshi{pixel.NewSprite(pic, pic.Bounds()), pixel.V(500, 500)}
}

func (y *Yoshi) Update() {
	velocity += gravity
	velocity *= 0.95
	y.V = y.V.Sub(pixel.V(0, velocity))
}

func (y *Yoshi) Lift() {
	velocity -= lift
}
func (y Yoshi) GetVel() pixel.Vec {
	return y.V
}

func (y Yoshi) TrueVel() float64 {
	return velocity
}

func (y Yoshi) GetSprite() *pixel.Sprite {
	return y.Sp
}
