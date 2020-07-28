package prophotorgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"math"
)

const PrimaryRedX = 0.734699
const PrimaryRedY = 0.265301
const PrimaryGreenX = 0.159597
const PrimaryGreenY = 0.840403
const PrimaryBlueX = 0.036598
const PrimaryBlueY = 0.000105

var StandardWhitePoint = ciexyz.WhitePointD50

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
