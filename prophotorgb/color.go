package prophotorgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/linear"
	"image/color"
)

// Color represents a linear normalised colour in Pro Photo RGB space.
type Color struct {
	linear.RGBColor
}

// ToNRGBA returns an encoded 8-bit NRGBA representation of this colour suitable
// for use with instances of image.NRGBA.
//
// alpha is the normalised alpha value and will be clipped to 0.0–1.0.
func (c Color) ToNRGBA(alpha float32) color.NRGBA {
	return c.RGBColor.ToNRGBA(alpha, To8Bit)
}

// ToRGBA returns an encoded 8-bit RGBA representation of this colour suitable
// for use with instances of image.RGBA.
//
// alpha is the normalised alpha value and will be clipped to 0.0–1.0.
func (c Color) ToRGBA(alpha float32) color.RGBA {
	return c.RGBColor.ToRGBA(alpha, To8Bit)
}

// ToXYZ returns a CIE XYZ representation of this colour.
func (c Color) ToXYZ() ciexyz.Color {
	return ciexyz.Color{
		X: c.R*0.7976734029450476 + c.G*0.13518768157360342 + c.B*0.03135091585137471,
		Y: c.R*0.28804113269427045 + c.G*0.7118689212431865 + c.B*8.994606254312391e-05,
		Z: c.R*0 + c.G*0 + c.B*0.8251882791519165,
	}
}

// ColorFromLinear creates a Color instance from a linear normalised RGB
// triplet.
func ColorFromLinear(r, g, b float32) Color {
	return Color{linear.RGBColor{R: r, G: g, B: b}}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Pro Photo RGB encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func ColorFromNRGBA(c color.NRGBA) (col Color, alpha float32) {
	rgb, a := linear.RGBColorFromNRGBA(c, From8Bit)
	return Color{rgb}, a
}

// ColorFromRGBA creates a Color instance by interpreting an 8-bit RGBA colour
// as Pro Photo RGB encoded. The alpha value is returned as a normalised value
// between 0.0–1.0.
func ColorFromRGBA(c color.RGBA) (col Color, alpha float32) {
	rgb, a := linear.RGBColorFromRGBA(c, From8Bit)
	return Color{rgb}, a
}

// ColorFromXYZ creates a Pro Photo RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return ColorFromLinear(
		c.X*1.3459441751995702+c.Y*-0.255601933365717+c.Z*-0.05110784199842092,
		c.X*-0.5446048748563069+c.Y*1.5081763487167865+c.Z*0.020526475602742393,
		c.X*0+c.Y*-0+c.Z*1.2118446483846639,
	)
}
