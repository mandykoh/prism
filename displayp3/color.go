package displayp3

import (
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/colconv"
	"github.com/mandykoh/prism/srgb"
	"image/color"
	"math"
)

// Color represents a linear normalised colour in Display P3 space.
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
		R: srgb.To8Bit(c.R),
		G: srgb.To8Bit(c.G),
		B: srgb.To8Bit(c.B),
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
		X: c.R*0.48656856264244125 + c.G*0.2656727168458704 + c.B*0.19818726598669462,
		Y: c.R*0.22897344124350177 + c.G*0.6917516599220641 + c.B*0.07927489883443435,
		Z: c.R*0 + c.G*0.04511425370425419 + c.B*1.0437861931875305,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Display P3 encoded. The alpha value is returned as a normalised
// value between 0.0–1.0.
func ColorFromNRGBA(c color.NRGBA) (col Color, alpha float32) {
	return Color{
			R: srgb.From8Bit(c.R),
			G: srgb.From8Bit(c.G),
			B: srgb.From8Bit(c.B),
		},
		float32(c.A) / 255
}

// ColorFromRGBA creates a Color instance by interpreting an 8-bit RGBA colour
// as Display P3 encoded. The alpha value is returned as a normalised value
// between 0.0–1.0.
func ColorFromRGBA(c color.RGBA) (col Color, alpha float32) {
	if c.A == 0 {
		return Color{}, 0
	}

	return Color{
			R: srgb.From8Bit(uint8((uint32(c.R) * 255) / uint32(c.A))),
			G: srgb.From8Bit(uint8((uint32(c.G) * 255) / uint32(c.A))),
			B: srgb.From8Bit(uint8((uint32(c.B) * 255) / uint32(c.A))),
		},
		float32(c.A) / 255
}

// ColorFromXYZ creates a Display P3 Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*2.493509087331807 + c.Y*-0.931388074532663 + c.Z*-0.40271279318557973,
		G: c.X*-0.8294731994547587 + c.Y*1.7626305488413623 + c.Z*0.0236242511428412,
		B: c.X*0.03585127357050431 + c.Y*-0.07618395633732165 + c.Z*0.9570295296681479,
	}
}
