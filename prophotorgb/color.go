package prophotorgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/colconv"
	"image/color"
	"math"
)

// Color represents a linear normalised colour in Pro Photo RGB space.
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
		X: c.R*0.7976734029450476 + c.G*0.13518768157360342 + c.B*0.03135091585137471,
		Y: c.R*0.28804113269427045 + c.G*0.7118689212431865 + c.B*8.994606254312391e-05,
		Z: c.R*0 + c.G*0 + c.B*0.8251882791519165,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Pro Photo RGB encoded. The alpha value is returned as a normalised
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
// as Pro Photo RGB encoded. The alpha value is returned as a normalised value
// between 0.0–1.0.
func ColorFromRGBA(c color.RGBA) (col Color, alpha float32) {
	return ColorFromNRGBA(colconv.RGBAtoNRGBA(c))
}

// ColorFromXYZ creates a Pro Photo RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*1.3459441751995702 + c.Y*-0.255601933365717 + c.Z*-0.05110784199842092,
		G: c.X*-0.5446048748563069 + c.Y*1.5081763487167865 + c.Z*0.020526475602742393,
		B: c.X*0 + c.Y*-0 + c.Z*1.2118446483846639,
	}
}
