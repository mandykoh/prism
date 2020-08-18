package prophotorgb

import (
	"github.com/mandykoh/prism/ciexyy"
	"github.com/mandykoh/prism/linear"
	"image"
	"image/color"
	"image/draw"
	"math"
)

var PrimaryRed = ciexyy.Color{X: 0.734699, Y: 0.265301, YY: 1}
var PrimaryGreen = ciexyy.Color{X: 0.159597, Y: 0.840403, YY: 1}
var PrimaryBlue = ciexyy.Color{X: 0.036598, Y: 0.000105, YY: 1}
var StandardWhitePoint = ciexyy.D50

const constantE = 1.0 / 512.0

// EncodeColor converts a linear colour value to a Pro Photo RGB encoded one.
func EncodeColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromLinearColor(c)
	return col.ToRGBA64(alpha)
}

// EncodeImage converts an image with linear colour into a Pro Photo RGB encoded
// one.
//
// src is the linearised image to be encoded.
//
// dst is the image to write the result to, beginning at its origin.
//
// src and dst may be the same image.
func EncodeImage(dst draw.Image, src image.Image) {
	linear.TransformImageColor(dst, src, EncodeColor)
}

func encodedToLinear(v float32) float32 {
	if v < constantE*16 {
		return v / 16
	}
	return float32(math.Pow(float64(v), 1.8))
}

// LineariseColor converts a Pro Photo RGB encoded colour into a linear one.
func LineariseColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromEncodedColor(c)
	return col.ToLinearRGBA64(alpha)
}

// LineariseImage converts an image with Pro Photo RGB encoded colour to linear
// colour.
//
// src is the encoded image to be linearised.
//
// dst is the image to write the result to, beginning at its origin.
//
// src and dst may be the same image.
func LineariseImage(dst draw.Image, src image.Image) {
	linear.TransformImageColor(dst, src, LineariseColor)
}

func linearToEncoded(v float32) float32 {
	if v < 0 {
		return 0
	} else if v < constantE {
		return 16 * v
	} else if v < 1 {
		return float32(math.Pow(float64(v), 1.0/1.8))
	}

	return 1
}
