package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (i Image) Bounds() image.Rectangle {
	rand.Seed(time.Now().UnixNano())
	w := 200 + rand.Intn(55)
	h := 200 + rand.Intn(55)
	return image.Rect(0, 0, w, h)
}

func (i Image) At(x, y int) color.Color {
	rand.Seed(time.Now().UnixNano())
	div := rand.Intn(200) + 1
	v := uint8(((x + y) / div) % 255)
	return color.RGBA{v, v, 255, 255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
