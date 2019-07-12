package bird_anim

import (
	"encoding/json"
	"fmt"
	"github.com/JChouCode/flappygo/bird"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"io/ioutil"
	"math"
	"os"
)

func loadAnimationJson(imgPath string, jsonPath string) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {
	// open and load the spritesheet
	sheetFile, err := os.Open(imgPath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)
	height := sheet.Bounds().H()
	// width := sheet.Bounds().W()

	temp, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, nil, err
	}

	type Pos struct {
		X float64
		Y float64
		W float64
		H float64
	}
	var data map[string][]Pos
	err2 := json.Unmarshal(temp, &data)
	if err2 != nil {
		return nil, nil, err
	}
	// fmt.Println(data)
	anims = make(map[string][]pixel.Rect)
	// load the animation information, name and interval inside the spritesheet

	for k, v := range data {
		for _, i := range v {
			if _, ok := anims[k]; ok {
				anims[k] = append(anims[k], pixel.R(i.X, height+i.Y-i.H, i.X+i.W, height+i.Y))
			} else {
				anims[k] = []pixel.Rect{pixel.R(i.X, height+i.Y-i.H, i.X+i.W, height+i.Y)}
			}
		}
	}
	// fmt.Println(anims)
	return sheet, anims, nil
}

func loadAnimationBird(imgPath string) {
	sheetFile, _ := os.Open(imgPath)
	// if err != nil {
	// 	return err
	// }
	defer sheetFile.Close()
	sheetImg, _, _ := image.Decode(sheetFile)
	// if err != nil {
	// 	return err
	// }
	var sheet pixel.Picture
	sheet = pixel.PictureDataFromImage(sheetImg)
	anims = append(anims, sheet)
	// return nil
}

var anims []pixel.Picture
var counter = 0.0
var rate = 1.0 / 10

type BirdAnim struct {
	sprite *pixel.Sprite
	sheet  pixel.Picture
	anims  []pixel.Picture
}

func New() BirdAnim {
	loadAnimationBird("grumpybird/fly1.png")
	loadAnimationBird("grumpybird/fly2.png")
	loadAnimationBird("grumpybird/fly3.png")
	loadAnimationBird("grumpybird/fly4.png")
	fmt.Println(anims)
	return BirdAnim{pixel.NewSprite(nil, pixel.Rect{}), nil, anims}
}

func (ba *BirdAnim) Update(b bird.Bird, dt float64) {
	counter += dt
	i := int(math.Floor(counter / rate))
	ba.sheet = ba.anims[i%len(ba.anims)]
}

func (ba *BirdAnim) Draw(t pixel.Target, b bird.Bird) {
	ba.sprite.Set(ba.sheet, ba.sheet.Bounds())
	ba.sprite.Draw(t, pixel.IM.ScaledXY(pixel.ZV, pixel.V(
		b.GetBody().W()/ba.sprite.Picture().Bounds().W(),
		b.GetBody().H()/ba.sprite.Picture().Bounds().H(),
	)).
		Moved(b.GetBody().Center()))
}
