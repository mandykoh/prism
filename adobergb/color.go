package adobergb

import (
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/colconv"
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

// ToRGBA returns an encoded 8-bit RGBA representation of this colour suitable
// for use with instances of image.RGBA.
//
// alpha is the normalised alpha value and will be clipped to 0.0–1.0.
func (c Color) ToRGBA(alpha float32) color.RGBA {
	return colconv.NRGBAtoRGBA(c.ToNRGBA(alpha))
}

// ToXYZ returns a CIE XYZ representation of this colour.
func (c Color) ToXYZ() ciexyz.Color {
	return ciexyz.Color{
		X: c.R*0.5766680793281725 + c.G*0.1855619421659935 + c.B*0.18819852398084014,
		Y: c.R*0.29734448781899253 + c.G*0.6273761097678748 + c.B*0.07527940241313279,
		Z: c.R*0.027031317880049893 + c.G*0.07069030664147563 + c.B*0.9911788223702592,
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

// ColorFromRGBA creates a Color instance by interpreting an 8-bit RGBA colour
// as Adobe RGB (1998) encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func ColorFromRGBA(c color.RGBA) (col Color, alpha float32) {
	return ColorFromNRGBA(colconv.RGBAtoNRGBA(c))
}

// ColorFromXYZ creates an Adobe RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*2.0415913017647322 + c.Y*-0.5650078698012716 + c.Z*-0.34473195659062167,
		G: c.X*-0.9692242864995342 + c.Y*1.8759299885141114 + c.Z*0.04155424903337176,
		B: c.X*0.013446472278330708 + c.Y*-0.11838142234726094 + c.Z*1.01533754937275,
	}
}
