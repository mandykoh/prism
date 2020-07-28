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
	A float32
}

// To8Bit returns an encoded 8-bit NRGBA representation of this colour suitable
// for use with instances of image.NRGBA.
func (c Color) To8Bit() color.NRGBA {
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
		X: c.R*0.5767309 + c.G*0.1855540 + c.B*0.1881852,
		Y: c.R*0.2973769 + c.G*0.6273491 + c.B*0.0752741,
		Z: c.R*0.0270343 + c.G*0.0706872 + c.B*0.9911085,
		A: c.A,
	}
}

// ColorFromNRGBA creates a Color instance by interpreting an 8-bit NRGBA colour
// as Adobe RGB (1998) encoded.
func ColorFromNRGBA(c color.NRGBA) Color {
	return Color{
		R: From8Bit(c.R),
		G: From8Bit(c.G),
		B: From8Bit(c.B),
		A: float32(c.A) * 255,
	}
}

// ColorFromXYZ creates an Adobe RGB Color instance from a CIE XYZ colour.
func ColorFromXYZ(c ciexyz.Color) Color {
	return Color{
		R: c.X*2.0413690 + c.Y*-0.5649464 + c.Z*-0.3446944,
		G: c.X*-0.9692660 + c.Y*1.8760108 + c.Z*0.0415560,
		B: c.X*0.0134474 + c.Y*-0.1183897 + c.Z*1.0154096,
		A: c.A,
	}
}
