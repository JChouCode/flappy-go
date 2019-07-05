package main

import (
	"./pipe"
	"./yoshi"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	// "image"
	"fmt"
	_ "image/png"
	"math/rand"
	// "os"
	"time"
)

func remove(slice []pipe.Pipe, s int) []pipe.Pipe {
	return append(slice[:s], slice[s+1:]...)
}

func collide(yosh yoshi.Yoshi, pipe pipe.Pipe) bool {
	return 500 <= pipe.X1 && 500 >= pipe.X2 && yosh.GetVel().Y >= pipe.Y1 && yosh.GetVel().Y <= pipe.Y2
}
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "BigYoshi",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	yosh := yoshi.New("bigyoshi.png")

	sprite := yosh.GetSprite()

	// pipes := []pipe.Pipe{pipe.New(rand.Intn(100) + 100, "top")
	// , pipe.New(rand.Intn(100) + 100, "bottom")}
	pipes := make([]pipe.Pipe, 0)

	pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), "top"))
	pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), "bottom"))

	imd := imdraw.New(nil)

	last := time.Now()

	for !win.Closed() {
		imd.Reset()
		imd.Clear()

		win.Clear(colornames.Blanchedalmond)

		if win.JustPressed(pixelgl.KeySpace) {
			yosh.Lift()
		}

		dt := time.Since(last).Seconds()
		if dt > 1.2 {
			pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), "top"))
			pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), "bottom"))
			last = time.Now()
		}

		mat := pixel.IM
		mat = mat.Scaled(pixel.ZV, 0.5)
		mat = mat.Moved(yosh.GetVel())
		sprite.Draw(win, mat)
		yosh.Update()

		for index, element := range pipes {
			imd.Push(element.GetCoord1(), element.GetCoord2())
			imd.Rectangle(0)
			(&pipes[index]).Update()
		}

		if len(pipes) != 0 && pipes[0].OutBounds() {
			// fmt.Println(pipes)
			pipes = pipes[1:]
		}

		for _, element := range pipes {
			if collide(yosh, element) {
				fmt.Println("collide")
				imd.Color = colornames.Red
			}
		}

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
