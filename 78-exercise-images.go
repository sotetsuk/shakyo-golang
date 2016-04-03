package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.Width, im.Height)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x) % 255, uint8(y) % 255, uint8(x*y) % 255, 255}
}

func main() {
	m := Image{100, 200}
	pic.ShowImage(m)
}
