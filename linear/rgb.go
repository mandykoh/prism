package linear

import (
	"image/color"
	"math"
)

// RGB represents a linear normalised RGB colour value in an unspecified colour
// space.
type RGB struct {
	R float32
	G float32
	B float32
}

// Luminance returns the perceptual luminance of this colour.
func (c RGB) Luminance() float32 {
	return 0.2126*c.R + 0.7152*c.G + 0.0722*c.B
}

// ToEncodedNRGBA returns an encoded 8-bit NRGBA representation of this colour suitable
// for use with instances of image.NRGBA.
//
// alpha is the normalised alpha value and will be clipped to 0.0–1.0.
//
// trcEncode is a tonal response curve encoding function.
func (c RGB) ToEncodedNRGBA(alpha float32, trcEncode func(float32) uint8) color.NRGBA {
	return color.NRGBA{
		R: trcEncode(c.R),
		G: trcEncode(c.G),
		B: trcEncode(c.B),
		A: uint8(math.Max(math.Min(float64(alpha), 1), 0) * 255),
	}
}

// ToEncodedRGBA returns an encoded 8-bit RGBA representation of this colour suitable
// for use with instances of image.RGBA.
//
// alpha is the normalised alpha value and will be clipped to 0.0–1.0.
//
// trcEncode is a tonal response curve encoding function.
func (c RGB) ToEncodedRGBA(alpha float32, trcEncode func(float32) uint8) color.RGBA {
	return color.RGBA{
		R: trcEncode(c.R * alpha),
		G: trcEncode(c.G * alpha),
		B: trcEncode(c.B * alpha),
		A: uint8(math.Max(math.Min(float64(alpha), 1.0), 0.0) * 255),
	}
}

// RGBFromEncodedNRGBA creates a Color instance by interpreting an 8-bit NRGBA
// colour as tonal response encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func RGBFromEncodedNRGBA(c color.NRGBA, trcDecode func(uint8) float32) (col RGB, alpha float32) {
	return RGB{
			R: trcDecode(c.R),
			G: trcDecode(c.G),
			B: trcDecode(c.B),
		},
		float32(c.A) / 255
}

// RGBFromEncodedRGBA creates a Color instance by interpreting an 8-bit RGBA colour
// as sRGB encoded. The alpha value is returned as a normalised value between
// 0.0–1.0.
func RGBFromEncodedRGBA(c color.RGBA, trcDecode func(uint8) float32) (col RGB, alpha float32) {
	if c.A == 0 {
		return RGB{}, 0
	}

	alpha = float32(c.A) / 255

	return RGB{
			R: trcDecode(c.R) / alpha,
			G: trcDecode(c.G) / alpha,
			B: trcDecode(c.B) / alpha,
		},
		alpha
}
