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
		X: c.R*0.4124564011253347 + c.G*0.35757608771164573 + c.B*0.18043748186614458,
		Y: c.R*0.2126728318302507 + c.G*0.7151521754232915 + c.B*0.07217499274645783,
		Z: c.R*0.019333893802750045 + c.G*0.11919202923721521 + c.B*0.9503040711616949,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as sRGB encoded. The alpha value is returned as a normalised value between
// 0.0–1.0.
func ColorFromNRGBA(c color.NRGBA) (col Color, alpha float32) {
	return Color{
			R: From8Bit(c.R),
			G: From8Bit(c.G),
			B: From8Bit(c.B),
		},
		float32(c.A) / 255
}

// ColorFromXYZ creates an SRGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*3.2404544603802004 + c.Y*-1.5371386542829153 + c.Z*-0.4985314554431076,
		G: c.X*-0.9692660032150083 + c.Y*1.8760107926266538 + c.Z*0.041556016360319,
		B: c.X*0.05564343139092609 + c.Y*-0.20402591510006216 + c.Z*1.057225196427595,
	}
}
