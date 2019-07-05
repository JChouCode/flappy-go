package pipe

import (
	"github.com/faiface/pixel"
	// "image"
	_ "image/png"
	// "os"
)

// func loadPicture(path string) (pixel.Picture, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pixel.PictureDataFromImage(img), nil
// }

const velocity = 5

type Pipe struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func New(height float64, pos string) Pipe {
	if pos == "top" {
		return Pipe{1024, 768, 994, 768 - height}
	}
	if pos == "bottom" {
		return Pipe{1024, 0, 994, height}
	}
	return Pipe{0, 0, 0, 0}
}

func (p *Pipe) Update() {
	p.X1 -= velocity
	p.X2 -= velocity
}

func (p Pipe) GetCoord1() pixel.Vec {
	return pixel.V(p.X1, p.Y1)
}

func (p Pipe) GetCoord2() pixel.Vec {
	return pixel.V(p.X2, p.Y2)
}

func (p Pipe) OutBounds() bool {
	return p.X2 < 0
}
