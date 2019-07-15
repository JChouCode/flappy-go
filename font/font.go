package font

import (
	"encoding/json"
	// "fmt"
	// "github.com/JChouCode/flappygo/bird"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"io/ioutil"
	// "math"
	"os"
	"strconv"
	"strings"
)

const sheight = 512.0
const swidth = 672.0

var anims map[int]*pixel.Sprite
var err error

func loadFontJson(imgPath string, jsonPath string) (anims map[int]*pixel.Sprite, err error) {
	// open and load the spritesheet
	sheetFile, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, err
	}
	sheet := pixel.PictureDataFromImage(sheetImg)
	height := sheet.Bounds().H()

	temp, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	type Pos struct {
		X float64
		Y float64
		W float64
		H float64
	}
	var data map[string]Pos
	err2 := json.Unmarshal(temp, &data)
	if err2 != nil {
		return nil, err
	}

	anims = make(map[int]*pixel.Sprite)
	for k, v := range data {
		temp, _ := strconv.Atoi(k)
		anims[temp] = pixel.NewSprite(sheet, pixel.R(v.X, height+v.Y-v.H, v.X+v.W, height+v.Y))
	}
	return anims, nil
}

func init() {
	anims, err = loadFontJson("flappyfont.png", "flappyfont.json")
}

func DrawScore(t pixel.Target, score int) {
	// Turn score into array of ints
	s := strconv.Itoa(score)
	split := strings.Split(s, "")
	scoreArr := make([]int, 0)
	for _, v := range split {
		j, _ := strconv.Atoi(v)
		scoreArr = append(scoreArr, j)
	}

	scoreWidth := 0.0
	xStart := swidth / 2
	yStart := sheight * 0.9
	spacing := 3.0

	// Calculate score width
	for _, v := range scoreArr {
		scoreWidth += anims[v].Frame().W()
	}
	xStart -= scoreWidth / 2

	// Draw score
	for _, v := range scoreArr {
		anims[v].Draw(t, pixel.IM.Moved(pixel.V(xStart, yStart)))
		xStart += anims[v].Frame().W() + spacing
	}
}
