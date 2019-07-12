package pipe

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

var velocity = pixel.V(-200, 0)

var sheet pixel.Picture

type Pipe struct {
	sprite *pixel.Sprite
	body   pixel.Rect
	bot    bool
}

func init() {
	sheet, _ = loadPicture("pipe.png")
}

func New(height float64, bot bool) Pipe {
	if bot {
		return Pipe{pixel.NewSprite(sheet, sheet.Bounds()), pixel.R(1024, height-sheet.Bounds().H(), 1024+sheet.Bounds().W(), height), bot}
	}
	return Pipe{pixel.NewSprite(sheet, sheet.Bounds()), pixel.R(1024, 768-height, 1024+sheet.Bounds().W(), 768-height+sheet.Bounds().H()), bot}
}

func (p *Pipe) Update(dt float64) {
	p.body = p.body.Moved(velocity.Scaled(dt))
}

func (p *Pipe) Draw(t pixel.Target) {
	if p.bot {
		p.sprite.Draw(t, pixel.IM.Moved(p.body.Center()))
	} else {
		p.sprite.Draw(t, pixel.IM.ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(p.body.Center()))
	}
}

func (p Pipe) Offscreen() bool {
	return p.body.Min.X < 0
}

func (p Pipe) GetBody() pixel.Rect {
	return p.body
}
