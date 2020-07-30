package prophotorgb

import (
	"github.com/mandykoh/prism/ciexyy"
	"math"
)

var PrimaryRed = ciexyy.Color{X: 0.734699, Y: 0.265301, YY: 1}
var PrimaryGreen = ciexyy.Color{X: 0.159597, Y: 0.840403, YY: 1}
var PrimaryBlue = ciexyy.Color{X: 0.036598, Y: 0.000105, YY: 1}
var StandardWhitePoint = ciexyy.D50

const constantE = 1.0 / 512.0

// ConvertLinearTo8Bit converts a linear value to an 8-bit ProPhoto RGB encoded
// value, clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see To8Bit.
func ConvertLinearTo8Bit(v float64) uint8 {
	var scaled float64
	if v < 0 {
		scaled = 0
	} else if v < constantE {
		scaled = 16 * v
	} else if v < 1 {
		scaled = math.Pow(v, 1.0/1.8)
	} else {
		scaled = 1
	}

	return uint8(math.Round(scaled * 255))
}

// Convert8BitToLinear converts an 8-bit ProPhoto RGB encoded value to a
// normalised linear value between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see From8Bit.
func Convert8BitToLinear(v uint8) float64 {
	vNormalised := float64(v) / 255
	if vNormalised < constantE*16 {
		return vNormalised / 16
	}
	return math.Pow(vNormalised, 1.8)
}
