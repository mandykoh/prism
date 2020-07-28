package ciexyz

import (
	"github.com/mandykoh/prism/cielab"
	"math"
)

const constantE = 216.0 / 24389.0
const constantK = 24389.0 / 27.0

// Color represents a linear normalised colour in CIE XYZ space.
type Color struct {
	X float32
	Y float32
	Z float32
	A float32
}

func (c Color) ToLAB(whitePoint Color) cielab.Color {
	fx := componentToLAB(c.X, whitePoint.X)
	fy := componentToLAB(c.Y, whitePoint.Y)
	fz := componentToLAB(c.Z, whitePoint.Z)

	return cielab.Color{
		L:     float32(116*fy - 16),
		A:     float32(500 * (fx - fy)),
		B:     float32(200 * (fy - fz)),
		Alpha: c.A,
	}
}

func ColorFromLAB(lab cielab.Color, whitePoint Color) Color {
	fy := (float64(lab.L) + 16) / 116
	fx := float64(lab.A)/500 + fy
	fz := fy - float64(lab.B)/200

	var xr float64
	if fx3 := math.Pow(fx, 3); fx3 > constantE {
		xr = fx3
	} else {
		xr = (116*fx - 16) / constantK
	}

	var yr float64
	if lab.L > constantK*constantE {
		yr = math.Pow((float64(lab.L)+16)/116, 3)
	} else {
		yr = float64(lab.L) / constantK
	}

	var zr float64
	if fz3 := math.Pow(fz, 3); fz3 > constantE {
		zr = fz3
	} else {
		zr = (116*fz - 16) / constantK
	}

	return Color{
		X: float32(xr * float64(whitePoint.X)),
		Y: float32(yr * float64(whitePoint.Y)),
		Z: float32(zr * float64(whitePoint.Z)),
		A: lab.Alpha,
	}
}

func componentToLAB(v float32, wp float32) float64 {
	r := float64(v) / float64(wp)
	if r > constantE {
		return math.Pow(r, 1.0/3.0)
	}
	return (constantK*r + 16) / 116.0
}
