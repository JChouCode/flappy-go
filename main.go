package main

import (
	// "fmt"
	"github.com/JChouCode/flappygo/bird"
	"github.com/JChouCode/flappygo/bird_anim"
	"github.com/JChouCode/flappygo/pipe"
	"github.com/JChouCode/flappygo/scrollable"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	_ "image/png"
	"math/rand"
	"time"
)

const sheight = 512.0
const swidth = 672.0

func collide(player bird.Bird, pipe pipe.Pipe) bool {
	// if player.GetBody().Min.X >= pipe.GetBody().Min.X &&
	// 	player.GetBody().Max.X <= pipe.GetBody().Max.X {
	// 	fmt.Println("collide x")
	// }
	// if player.GetBody().Min.Y >= pipe.GetBody().Min.Y &&
	// 	player.GetBody().Max.Y <= pipe.GetBody().Max.Y {
	// 	fmt.Println("collide y")
	// }
	return swidth*0.4+player.GetBody().W() >= pipe.GetBody().Min.X &&
		swidth*0.4+player.GetBody().W() <= pipe.GetBody().Max.X &&
		player.GetBody().Min.Y >= pipe.GetBody().Min.Y-20 &&
		player.GetBody().Max.Y <= pipe.GetBody().Max.Y+20
	// return false
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
	// win.SetSmooth(true)
	return win
}

func run() {
	win := initWindow("Flappy-GO", 672, 512)

	player := bird.New()
	anim := bird_anim.New()

	var pipes []pipe.Pipe
	pipeBot, pipeTop := pipe.New()
	pipes = append(pipes, pipeTop)
	pipes = append(pipes, pipeBot)

	bases := scrollable.New("base.png", 112, 336)
	background := scrollable.New("background-day.png", 512, 288)

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
			pipeBot, pipeTop := pipe.New()
			pipes = append(pipes, pipeTop)
			pipes = append(pipes, pipeBot)
			pipeLast = time.Now()
		}

		//Update background
		for i := range background {
			(&background[i]).Update(dt)
		}
		player.Update(dt)
		anim.Update(player, dt)
		// Update pipes
		for i := range pipes {
			(&pipes[i]).Update(dt)
		}

		//Update bases
		for i := range bases {
			(&bases[i]).Update(dt)
		}

		// Remove offscreen pipes
		k := 0
		for _, pipe := range pipes {
			if !pipe.Offscreen() {
				pipes[k] = pipe
				k++
			}
		}
		pipes = pipes[:k]

		//Cycle offscreen bases
		if bases[0].GetBody().Max.X < 0 {
			bases = scrollable.Cycle(bases)
		}

		//Cycle offscreen background
		if background[0].GetBody().Max.X < 0 {
			background = scrollable.Cycle(background)
		}

		for _, element := range pipes {
			if collide(player, element) {
				player.Reset()
				pipes = nil
			}
		}

		for _, b := range background {
			b.Draw(win)
		}
		anim.Draw(win, player)
		for _, pipe := range pipes {
			pipe.Draw(win)
		}
		for _, b := range bases {
			b.Draw(win)
		}

		//Debug
		// fmt.Fprintln(basicTxt, player.GetBody())
		// fmt.Println(player.GetBody())
		// fmt.Fprintln(basicTxt, len(pipes))
		// if len(pipes) > 0 {
		// 	fmt.Fprintln(basicTxt, pipes[0].GetBody())
		// }
		// fmt.Println(pipes[0].GetBody())
		// fmt.Fprintln(basicTxt, pipe.)

		basicTxt.Draw(win, pixel.IM)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
