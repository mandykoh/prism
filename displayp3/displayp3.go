package displayp3

import (
	"github.com/mandykoh/prism/ciexyy"
	"image"
	"image/color"
	"image/draw"
)

var PrimaryRed = ciexyy.Color{X: 0.68, Y: 0.32, YY: 1}
var PrimaryGreen = ciexyy.Color{X: 0.265, Y: 0.69, YY: 1}
var PrimaryBlue = ciexyy.Color{X: 0.15, Y: 0.06, YY: 1}
var StandardWhitePoint = ciexyy.D65

// EncodeColor converts a linear colour value to a Display P3 encoded one.
func EncodeColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromLinearColor(c)
	return col.ToRGBA64(alpha)
}

// EncodeImage converts an image with linear colour into a Display P3 encoded
// one.
func EncodeImage(img draw.Image) {
	bounds := img.Bounds()

	switch inputImg := img.(type) {

	case *image.RGBA64:
		for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
			for j := bounds.Min.X; j < bounds.Max.X; j++ {
				inputImg.SetRGBA64(j, i, EncodeColor(inputImg.RGBA64At(j, i)))
			}
		}

	default:
		for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
			for j := bounds.Min.X; j < bounds.Max.X; j++ {
				img.Set(j, i, EncodeColor(img.At(j, i)))
			}
		}
	}
}

// LineariseColor converts a Display P3 encoded colour into a linear one.
func LineariseColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromEncodedColor(c)
	return col.ToLinearRGBA64(alpha)
}

// LineariseImage converts an image with Display P3 encoded colour to linear
// colour.
func LineariseImage(img draw.Image) {
	bounds := img.Bounds()

	switch inputImg := img.(type) {

	case *image.RGBA64:
		for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
			for j := bounds.Min.X; j < bounds.Max.X; j++ {
				inputImg.SetRGBA64(j, i, LineariseColor(inputImg.RGBA64At(j, i)))
			}
		}

	default:
		for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
			for j := bounds.Min.X; j < bounds.Max.X; j++ {
				img.Set(j, i, LineariseColor(img.At(j, i)))
			}
		}
	}
}
