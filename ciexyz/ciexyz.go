package ciexyz

import (
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

func invertMatrix(m [3][3]float64) [3][3]float64 {
	o := [3][3]float64{
		{
			m[1][1]*m[2][2] - m[2][1]*m[1][2],
			-(m[0][1]*m[2][2] - m[2][1]*m[0][2]),
			m[0][1]*m[1][2] - m[1][1]*m[0][2],
		},
		{
			-(m[1][0]*m[2][2] - m[2][0]*m[1][2]),
			m[0][0]*m[2][2] - m[2][0]*m[0][2],
			-(m[0][0]*m[1][2] - m[1][0]*m[0][2]),
		},
		{
			m[1][0]*m[2][1] - m[2][0]*m[1][1],
			-(m[0][0]*m[2][1] - m[2][0]*m[0][1]),
			m[0][0]*m[1][1] - m[1][0]*m[0][1],
		},
	}

	det := m[0][0]*o[0][0] + m[1][0]*o[0][1] + m[2][0]*o[0][2]
	if det == 0 {
		panic("matrix is non-invertible")
	}

	o[0][0] /= det
	o[0][1] /= det
	o[0][2] /= det
	o[1][0] /= det
	o[1][1] /= det
	o[1][2] /= det
	o[2][0] /= det
	o[2][1] /= det
	o[2][2] /= det

	return o
}

// TransformFromXYZForPrimaries generates the column matrix for converting
// colour values from CIE XYZ to a space defined by three RGB primary
// chromaticities and a reference white.
func TransformFromXYZForPrimaries(rx, ry, gx, gy, bx, by float64, whitePoint Color) [3][3]float64 {
	t := TransformToXYZForPrimaries(rx, ry, gx, gy, bx, by, whitePoint)
	return invertMatrix(t)
}

// TransformToXYZForPrimaries generates the column matrix for converting colour
// values from a space defined by three primary RGB chromaticities and a
// reference white to CIE XYZ.
func TransformToXYZForPrimaries(rx, ry, gx, gy, bx, by float64, whitePoint Color) [3][3]float64 {
	m := [3][3]float64{
		transformForComponent(rx, ry),
		transformForComponent(gx, gy),
		transformForComponent(bx, by),
	}

	i := invertMatrix(m)

	sr := float64(whitePoint.X)*i[0][0] + float64(whitePoint.Y)*i[1][0] + float64(whitePoint.Z)*i[2][0]
	sg := float64(whitePoint.X)*i[0][1] + float64(whitePoint.Y)*i[1][1] + float64(whitePoint.Z)*i[2][1]
	sb := float64(whitePoint.X)*i[0][2] + float64(whitePoint.Y)*i[1][2] + float64(whitePoint.Z)*i[2][2]

	return [3][3]float64{
		{sr * m[0][0], sr * m[0][1], sr * m[0][2]},
		{sg * m[1][0], sg * m[1][1], sg * m[1][2]},
		{sb * m[2][0], sb * m[2][1], sb * m[2][2]},
	}
}

func transformForComponent(vx, vy float64) [3]float64 {
	return [...]float64{
		vx / vy,
		1.0,
		(1 - vx - vy) / vy,
	}
}
