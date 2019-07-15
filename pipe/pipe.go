package pipe

import (
	// "fmt"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"math/rand"
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
const space = 70.0

var velocity = pixel.V(-200, 0)

var sheet pixel.Picture

type Pipe struct {
	sprite *pixel.Sprite
	body   pixel.Rect
	bot    bool
	passed bool
}

func init() {
	sheet, _ = loadPicture("pipe.png")
}

func New() (Pipe, Pipe) {
	heightBot, heightTop := CalculateNewPipe()
	return Pipe{pixel.NewSprite(sheet, sheet.Bounds()), pixel.R(swidth, heightBot-sheet.Bounds().H(), swidth+sheet.Bounds().W(), heightBot), true, false},
		Pipe{pixel.NewSprite(sheet, sheet.Bounds()), pixel.R(swidth, sheight-heightTop, swidth+sheet.Bounds().W(), sheight-heightTop+sheet.Bounds().H()), false, false}
}

func (p *Pipe) Update(dt float64) {
	p.body = p.body.Moved(velocity.Scaled(dt))
}

func (p *Pipe) Passed() {
	p.passed = true
}

func (p Pipe) GetPassed() bool {
	return p.passed
}

func (p Pipe) GetBot() bool {
	return p.bot
}

func (p *Pipe) Draw(t pixel.Target) {
	if p.bot {
		p.sprite.Draw(t, pixel.IM.Moved(p.body.Center()))
	} else {
		p.sprite.Draw(t, pixel.IM.ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(p.body.Center()))
	}
}

func (p Pipe) Offscreen() bool {
	return p.body.Max.X < 0
}

func (p Pipe) GetBody() pixel.Rect {
	return p.body
}

func CalculateNewPipe() (float64, float64) {
	start := float64(rand.Intn(sheight - 340))
	heightBot := 200 + start
	heightTop := sheight - heightBot - space
	return heightBot, heightTop
}
