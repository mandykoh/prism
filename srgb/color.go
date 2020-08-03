package srgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/colconv"
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
		X: c.R*0.41238652374145657 + c.G*0.35759149384555927 + c.B*0.18045052788799043,
		Y: c.R*0.21263680803732615 + c.G*0.7151829876911185 + c.B*0.07218020427155547,
		Z: c.R*0.019330619488581533 + c.G*0.11919711488225383 + c.B*0.9503727125209493,
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

// ColorFromRGBA creates a Color instance by interpreting an 8-bit RGBA colour
// as sRGB encoded. The alpha value is returned as a normalised value between
// 0.0–1.0.
func ColorFromRGBA(c color.RGBA) (col Color, alpha float32) {
	return ColorFromNRGBA(colconv.RGBAtoNRGBA(c))
}

// ColorFromXYZ creates an SRGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*3.241003600540255 + c.Y*-1.5373991710891957 + c.Z*-0.49861598312439226,
		G: c.X*-0.9692242864995344 + c.Y*1.8759299885141119 + c.Z*0.04155424903337176,
		B: c.X*0.05563936186796137 + c.Y*-0.20401108051523723 + c.Z*1.0571488385644063,
	}
}
