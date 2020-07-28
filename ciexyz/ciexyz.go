package ciexyz

import (
	"github.com/mandykoh/prism/matrix"
	"math"
)

var WhitePointD50 = Color{0.9642, 1.0, 0.8251, 1.0}
var WhitePointD65 = Color{0.95047, 1.0, 1.08883, 1.0}

func componentFromLAB(f float64) float64 {
	if f3 := math.Pow(f, 3); f3 > constantE {
		return f3
	}
	return (116*f - 16) / constantK
}

func componentToLAB(v float32, wp float32) float64 {
	r := float64(v) / float64(wp)
	if r > constantE {
		return math.Pow(r, 1.0/3.0)
	}
	return (constantK*r + 16) / 116.0
}

// TransformFromXYZForPrimaries generates the column matrix for converting
// colour values from CIE XYZ to a space defined by three RGB primary
// chromaticities and a reference white.
func TransformFromXYZForPrimaries(rx, ry, gx, gy, bx, by float64, whitePoint Color) matrix.Matrix3 {
	return TransformToXYZForPrimaries(rx, ry, gx, gy, bx, by, whitePoint).Inverse()
}

// TransformToXYZForPrimaries generates the column matrix for converting colour
// values from a space defined by three primary RGB chromaticities and a
// reference white to CIE XYZ.
func TransformToXYZForPrimaries(rx, ry, gx, gy, bx, by float64, whitePoint Color) matrix.Matrix3 {
	m := matrix.Matrix3{
		transformForComponent(rx, ry),
		transformForComponent(gx, gy),
		transformForComponent(bx, by),
	}

	s := m.Inverse().MulV(matrix.Vector3{
		float64(whitePoint.X),
		float64(whitePoint.Y),
		float64(whitePoint.Z),
	})

	return matrix.Matrix3{
		m[0].MulS(s[0]),
		m[1].MulS(s[1]),
		m[2].MulS(s[2]),
	}
}

func transformForComponent(vx, vy float64) matrix.Vector3 {
	return matrix.Vector3{
		vx / vy,
		1.0,
		(1 - vx - vy) / vy,
	}
}
