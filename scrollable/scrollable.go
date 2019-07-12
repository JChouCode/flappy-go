package scrollable

import (
	// "fmt"
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

const sheight = 512.0
const swidth = 672.0

var velocity = pixel.V(-200, 0)

type Scrollable struct {
	sprite *pixel.Sprite
	body   pixel.Rect
}

func New(imgPath string, height float64, width float64) []Scrollable {
	pic, _ := loadPicture(imgPath)
	var scrollable []Scrollable
	scrollable = append(scrollable, Scrollable{pixel.NewSprite(pic, pixel.R(0, 0, pic.Bounds().W(), pic.Bounds().H())), pixel.R(0, 0, width, height)})
	scrollable = append(scrollable, Scrollable{pixel.NewSprite(pic, pixel.R(0, 0, pic.Bounds().W(), pic.Bounds().H())), pixel.R(width, 0, width*2, height)})
	scrollable = append(scrollable, Scrollable{pixel.NewSprite(pic, pixel.R(0, 0, pic.Bounds().W(), pic.Bounds().H())), pixel.R(width*2, 0, width*3, height)})
	scrollable = append(scrollable, Scrollable{pixel.NewSprite(pic, pixel.R(0, 0, pic.Bounds().W(), pic.Bounds().H())), pixel.R(width*3, 0, width*4, height)})
	return scrollable
}

func (b *Scrollable) Update(dt float64) {
	b.body = b.body.Moved(velocity.Scaled(dt))
}

func (b *Scrollable) Draw(t pixel.Target) {
	b.sprite.Draw(t, pixel.IM.Moved(b.body.Center()))
}

func (b Scrollable) GetBody() pixel.Rect {
	return b.body
}

func Cycle(scrollable []Scrollable) []Scrollable {
	temp := scrollable[0]
	temp.body.Max.X = scrollable[len(scrollable)-1].body.Max.X + temp.body.W()
	temp.body.Min.X = scrollable[len(scrollable)-1].body.Max.X
	scrollable = scrollable[1:]
	scrollable = append(scrollable, temp)
	return scrollable
}
