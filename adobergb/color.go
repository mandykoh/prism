package adobergb

import (
	"github.com/mandykoh/prism/ciexyz"
	"image/color"
	"math"
)

// Color represents a linear normalised colour in Adobe RGB (1998) space.
type Color struct {
	R float32
	G float32
	B float32
}

// Luminance returns the perceptual luminance of this colour.
func (c Color) Luminance() float32 {
	return 0.2126*c.R + 0.7152*c.G + 0.0722*c.B
}

// ToNRGBA returns an encoded 8-bit NRGBA representation of this colour suitable
// for use with instances of image.NRGBA.
//
// alpha is the normalised alpha value and will be clipped to 0.0–1.0.
func (c Color) ToNRGBA(alpha float32) color.NRGBA {
	return color.NRGBA{
		R: To8Bit(c.R),
		G: To8Bit(c.G),
		B: To8Bit(c.B),
		A: uint8(math.Max(math.Min(float64(alpha), 1), 0) * 255),
	}
}

// ToXYZ returns a CIE XYZ representation of this colour.
func (c Color) ToXYZ() ciexyz.Color {
	return ciexyz.Color{
		X: c.R*0.5767308538590208 + c.G*0.1855539559355801 + c.B*0.18818516090852389,
		Y: c.R*0.29737684652105756 + c.G*0.6273490891155328 + c.B*0.07527406436340955,
		Z: c.R*0.027034258774641568 + c.G*0.07068722130879249 + c.B*0.9911085141182259,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Adobe RGB (1998) encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func ColorFromNRGBA(c color.NRGBA) (col Color, alpha float32) {
	return Color{
			R: From8Bit(c.R),
			G: From8Bit(c.G),
			B: From8Bit(c.B),
		},
		float32(c.A) / 255
}

// ColorFromXYZ creates an Adobe RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*2.0413690972656604 + c.Y*-0.5649464198330968 + c.Z*-0.3446944043036243,
		G: c.X*-0.969266003215008 + c.Y*1.8760107926266532 + c.Z*0.04155601636031898,
		B: c.X*0.013447387300642133 + c.Y*-0.11838974309780925 + c.Z*1.0154095783288708,
	}
}
