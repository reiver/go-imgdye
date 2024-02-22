package imgdye

import (
	"image"
	"image/color"
	"image/draw"
)

// Dye colors the whole image one single color.
func Dye(img draw.Image, c color.Color) {
	var uniform image.Uniform = image.Uniform{c}
	var background image.Image = &uniform

	var drawer draw.Drawer = draw.Src
	drawer.Draw(img, img.Bounds(), background, image.ZP)
}
