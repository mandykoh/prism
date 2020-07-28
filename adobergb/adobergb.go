package adobergb

import (
	"github.com/mandykoh/prism/ciexyz"
	"math"
)

const PrimaryRedX = 0.64
const PrimaryRedY = 0.33
const PrimaryGreenX = 0.21
const PrimaryGreenY = 0.71
const PrimaryBlueX = 0.15
const PrimaryBlueY = 0.06

var StandardWhitePoint = ciexyz.WhitePointD65

// ConvertLinearTo8Bit converts a linear value to an 8-bit Adobe RGB encoded
// value, clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see To8Bit.
func ConvertLinearTo8Bit(v float64) uint8 {
	scaled := math.Pow(v, 256.0/563)
	return uint8(math.Round(math.Min(math.Max(scaled, 0.0), 1.0) * 255))
}

// Convert8BitToLinear converts an 8-bit Adobe RGB encoded value to a normalised
// linear value between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see From8Bit.
func Convert8BitToLinear(v uint8) float64 {
	return math.Pow(float64(v)/255, 563.0/256)
}
