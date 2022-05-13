package tour

import (
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

//ColorModel should return color.RGBAModel.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

//At should return a color; the value v in the last picture generator
//corresponds to color.RGBA{v, v, 255, 255} in this one.
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}
