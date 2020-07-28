package prophotorgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"image/color"
	"math"
)

// Color represents a linear normalised colour in Pro Photo RGB space.
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

// Luminance returns the perceptual luminance of this colour.
func (c Color) Luminance() float32 {
	return 0.2126*c.R + 0.7152*c.G + 0.0722*c.B
}

// ToNRGBA returns an encoded 8-bit NRGBA representation of this colour suitable
// for use with instances of image.NRGBA.
func (c Color) ToNRGBA() color.NRGBA {
	return color.NRGBA{
		R: To8Bit(c.R),
		G: To8Bit(c.G),
		B: To8Bit(c.B),
		A: uint8(math.Max(math.Min(float64(c.A), 1), 0) * 255),
	}
}

// ToXYZ returns a CIE XYZ representation of this colour.
func (c Color) ToXYZ() ciexyz.Color {
	return ciexyz.Color{
		X: c.R*0.7976641331391671 + c.G*0.13518832725191376 + c.B*0.031347559445345105,
		Y: c.R*0.2880378116561397 + c.G*0.711872251906302 + c.B*8.99364375583703e-05,
		Z: c.R*0 + c.G*0 + c.B*0.8251000046730042,
		A: c.A,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Pro Photo RGB encoded.
func ColorFromNRGBA(c color.NRGBA) Color {
	return Color{
		R: From8Bit(c.R),
		G: From8Bit(c.G),
		B: From8Bit(c.B),
		A: float32(c.A) / 255,
	}
}

// ColorFromXYZ creates a Pro Photo RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*1.3459598353730142 + c.Y*-0.25560493221231595 + c.Z*-0.05110843232886563,
		G: c.X*-0.5446023840931071 + c.Y*1.5081693133113774 + c.Z*0.02052638000029258,
		B: c.X*0 + c.Y*-0 + c.Z*1.2119742992806193,
		A: c.A,
	}
}
