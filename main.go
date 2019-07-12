package main

import (
	"fmt"
	"github.com/JChouCode/flappygo/bird"
	"github.com/JChouCode/flappygo/bird_anim"
	"github.com/JChouCode/flappygo/pipe"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	_ "image/png"
	"math/rand"
	"time"
)

func collide(player bird.Bird, pipe pipe.Pipe) bool {
	if player.GetBody().Min.X >= pipe.GetBody().Min.X &&
		player.GetBody().Max.X <= pipe.GetBody().Max.X {
		fmt.Println("collide x")
	}
	if player.GetBody().Min.Y >= pipe.GetBody().Min.Y &&
		player.GetBody().Max.Y <= pipe.GetBody().Max.Y {
		fmt.Println("collide y")
	}
	return 500 >= pipe.GetBody().Min.X &&
		500 <= pipe.GetBody().Max.X &&
		player.GetBody().Min.Y >= pipe.GetBody().Min.Y &&
		player.GetBody().Max.Y <= pipe.GetBody().Max.Y
}

//Initialize window
func initWindow(t string, w float64, h float64) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  t,
		Bounds: pixel.R(0, 0, w, h),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)
	return win
}

func run() {
	win := initWindow("Flappy-GO", 1024, 768)

	player := bird.New()
	anim := bird_anim.New()

	var pipes []pipe.Pipe
	pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), false))
	pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), true))
	last := time.Now()
	pipeLast := time.Now()

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(200, 400), atlas)
	basicTxt.Color = colornames.Black

	for !win.Closed() {
		rand.Seed(time.Now().UnixNano())
		dt := time.Since(last).Seconds()
		last = time.Now()
		pipeDt := time.Since(pipeLast).Seconds()

		win.Clear(colornames.Blanchedalmond)
		basicTxt.Clear()

		if win.JustPressed(pixelgl.KeySpace) {
			player.Lift()
		}

		if pipeDt > 1.2 {
			pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), false))
			pipes = append(pipes, pipe.New(float64(rand.Intn(100)+200), true))
			pipeLast = time.Now()
		}

		player.Update(dt)
		anim.Update(player, dt)
		// Update pipes
		for i := range pipes {
			(&pipes[i]).Update(dt)
		}

		k := 0
		for _, pipe := range pipes {
			if !pipe.Offscreen() {
				pipes[k] = pipe
				k++
			}
		}
		pipes = pipes[:k]

		for _, element := range pipes {
			if collide(player, element) {
				player.Reset()
				pipes = nil
			}
		}

		anim.Draw(win, player)
		for _, pipe := range pipes {
			pipe.Draw(win)
		}

		//Debug
		fmt.Fprintln(basicTxt, player.GetBody())
		// fmt.Println(player.GetBody())
		fmt.Fprintln(basicTxt, len(pipes))
		if len(pipes) > 0 {
			fmt.Fprintln(basicTxt, pipes[0].GetBody())
		}
		// fmt.Println(pipes[0].GetBody())
		// fmt.Fprintln(basicTxt, pipe.)

		basicTxt.Draw(win, pixel.IM)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
