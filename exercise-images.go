package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model { return color.RGBAModel }

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	j := uint8(x * y)
	return color.RGBA{j, 2 * j, 3 * j, j}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
