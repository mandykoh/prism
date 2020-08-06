package adobergb

import (
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/linear"
	"image/color"
)

// Color represents a linear normalised colour in Adobe RGB (1998) space.
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
		X: c.R*0.5766680793281725 + c.G*0.1855619421659935 + c.B*0.18819852398084014,
		Y: c.R*0.29734448781899253 + c.G*0.6273761097678748 + c.B*0.07527940241313279,
		Z: c.R*0.027031317880049893 + c.G*0.07069030664147563 + c.B*0.9911788223702592,
	}
}

// ColorFromLinear creates a Color instance from a linear normalised RGB
// triplet.
func ColorFromLinear(r, g, b float32) Color {
	return Color{linear.RGBColor{R: r, G: g, B: b}}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Adobe RGB (1998) encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func ColorFromNRGBA(c color.NRGBA) (col Color, alpha float32) {
	rgb, a := linear.RGBColorFromNRGBA(c, From8Bit)
	return Color{rgb}, a
}

// ColorFromRGBA creates a Color instance by interpreting an 8-bit RGBA colour
// as Adobe RGB (1998) encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func ColorFromRGBA(c color.RGBA) (col Color, alpha float32) {
	rgb, a := linear.RGBColorFromRGBA(c, From8Bit)
	return Color{rgb}, a
}

// ColorFromXYZ creates an Adobe RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return ColorFromLinear(
		c.X*2.0415913017647322+c.Y*-0.5650078698012716+c.Z*-0.34473195659062167,
		c.X*-0.9692242864995342+c.Y*1.8759299885141114+c.Z*0.04155424903337176,
		c.X*0.013446472278330708+c.Y*-0.11838142234726094+c.Z*1.01533754937275,
	)
}
