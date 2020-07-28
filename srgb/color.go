package srgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"image/color"
	"math"
)

// Color represents a linear normalised colour in sRGB space.
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

// To8Bit returns an encoded 8-bit NRGBA representation of this colour suitable
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
		X: c.R*0.4124564 + c.G*0.3575761 + c.B*0.1804375,
		Y: c.R*0.2126729 + c.G*0.7151522 + c.B*0.0721750,
		Z: c.R*0.0193339 + c.G*0.1191920 + c.B*0.9503041,
		A: c.A,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as sRGB encoded.
func ColorFromNRGBA(c color.NRGBA) Color {
	return Color{
		R: From8Bit(c.R),
		G: From8Bit(c.G),
		B: From8Bit(c.B),
		A: float32(c.A) / 255,
	}
}

// ColorFromXYZ creates an SRGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*3.2404542 + c.Y*-1.5371385 + c.Z*-0.4985314,
		G: c.X*-0.9692660 + c.Y*1.8760108 + c.Z*0.0415560,
		B: c.X*0.0556434 + c.Y*-0.2040259 + c.Z*1.0572252,
		A: c.A,
	}
}
