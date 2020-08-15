package adobergb

import (
	"github.com/mandykoh/prism/ciexyy"
	"image"
	"image/color"
	"image/draw"
	"math"
)

var PrimaryRed = ciexyy.Color{X: 0.64, Y: 0.33, YY: 1}
var PrimaryGreen = ciexyy.Color{X: 0.21, Y: 0.71, YY: 1}
var PrimaryBlue = ciexyy.Color{X: 0.15, Y: 0.06, YY: 1}
var StandardWhitePoint = ciexyy.D65

// ConvertLinearTo8Bit converts a linear value to an 8-bit Adobe RGB encoded
// value, clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see To8Bit.
func ConvertLinearTo8Bit(v float32) uint8 {
	return uint8(math.Round(linearToEncoded(float64(v)) * 255))
}

// ConvertLinearTo16Bit converts a linear value to a 16-bit Adobe RGB encoded
// value, clipping the linear value to between 0.0 and 1.0.
func ConvertLinearTo16Bit(v float32) uint16 {
	return uint16(math.Round(linearToEncoded(float64(v)) * 65535))
}

// Convert8BitToLinear converts an 8-bit Adobe RGB encoded value to a normalised
// linear value between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see From8Bit.
func Convert8BitToLinear(v uint8) float32 {
	return float32(encodedToLinear(float64(v) / 255))
}

// Convert16BitToLinear converts a 16-bit Adobe RGB encoded value to a
// normalised linear value between 0.0 and 1.0.
func Convert16BitToLinear(v uint16) float32 {
	return float32(encodedToLinear(float64(v) / 65535))
}

// EncodeColor converts a linear colour value to an Adobe RGB encoded one.
func EncodeColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromLinearColor(c)
	return col.ToRGBA64(alpha)
}

// EncodeImage converts an image with linear colour into an Adobe RGB encoded
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

func encodedToLinear(v float64) float64 {
	return math.Pow(v, 563.0/256)
}

// LineariseColor converts an Adobe RGB encoded colour into a linear one.
func LineariseColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromEncodedColor(c)
	return col.ToLinearRGBA64(alpha)
}

// LineariseImage converts an image with Adobe RGB encoded colour to linear
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

func linearToEncoded(v float64) float64 {
	scaled := math.Pow(v, 256.0/563)
	return math.Min(math.Max(scaled, 0.0), 1.0)
}
