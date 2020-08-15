package srgb

import (
	"github.com/mandykoh/prism/linear"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/mandykoh/prism/ciexyy"
)

var PrimaryRed = ciexyy.Color{X: 0.64, Y: 0.33, YY: 1}
var PrimaryGreen = ciexyy.Color{X: 0.3, Y: 0.6, YY: 1}
var PrimaryBlue = ciexyy.Color{X: 0.15, Y: 0.06, YY: 1}
var StandardWhitePoint = ciexyy.D65

// ConvertLinearTo8Bit converts a linear value to an 8-bit sRGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see To8Bit.
func ConvertLinearTo8Bit(v float32) uint8 {
	return uint8(math.Round(linearToEncoded(float64(v)) * 255))
}

// ConvertLinearTo16Bit converts a linear value to a 16-bit sRGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
func ConvertLinearTo16Bit(v float32) uint16 {
	return uint16(math.Round(linearToEncoded(float64(v)) * 65535))
}

// Convert8BitToLinear converts an 8-bit sRGB encoded value to a normalised
// linear value between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see From8Bit.
func Convert8BitToLinear(v uint8) float32 {
	return float32(encodedToLinear(float64(v) / 255))
}

// Convert16BitToLinear converts a 16-bit sRGB encoded value to a normalised
// linear value between 0.0 and 1.0.
func Convert16BitToLinear(v uint16) float32 {
	return float32(encodedToLinear(float64(v) / 65535))
}

// EncodeColor converts a linear colour value to an sRGB encoded one.
func EncodeColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromLinearColor(c)
	return col.ToRGBA64(alpha)
}

// EncodeImage converts an image with linear colour into an sRGB encoded one.
//
// src is the linearised image to be encoded.
//
// dst is the image to write the result to, beginning at its origin.
//
// src and dst may be the same image.
func EncodeImage(dst draw.Image, src image.Image) {
	linear.TransformImageColor(dst, src, EncodeColor)
}

func encodedToLinear(v float64) float64 {
	if v <= 0.0031308*12.92 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

// LineariseColor converts an sRGB encoded colour into a linear one.
func LineariseColor(c color.Color) color.RGBA64 {
	col, alpha := ColorFromEncodedColor(c)
	return col.ToLinearRGBA64(alpha)
}

// LineariseImage converts an image with sRGB encoded colour to linear colour.
//
// src is the encoded image to be linearised.
//
// dst is the image to write the result to, beginning at its origin.
//
// src and dst may be the same image.
func LineariseImage(dst draw.Image, src image.Image) {
	linear.TransformImageColor(dst, src, LineariseColor)
}

func linearToEncoded(v float64) float64 {
	var scaled float64
	if v <= 0.0031308 {
		scaled = v * 12.92
	} else {
		scaled = 1.055*math.Pow(v, 1/2.4) - 0.055
	}
	return math.Min(math.Max(scaled, 0.0), 1.0)
}
