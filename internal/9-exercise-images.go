package main

import (
	"image"
	"image/color"	
	"golang.org/x/tour/pic"
)
type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle {
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 255, Y: 255},
	}
}

func (img Image) At(x, y int) color.Color {
    v := uint8(x|y)
	return color.RGBA{v, uint8(x+y), uint8(x-y), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
