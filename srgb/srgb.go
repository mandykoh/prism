package srgb

import (
	"github.com/mandykoh/prism/ciexyz"
	"math"
)

const PrimaryRedX = 0.64
const PrimaryRedY = 0.33
const PrimaryGreenX = 0.3
const PrimaryGreenY = 0.6
const PrimaryBlueX = 0.15
const PrimaryBlueY = 0.06

var StandardWhitePoint = ciexyz.WhitePointD65

// ConvertLinearTo8Bit converts a linear value to an 8-bit sRGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see To8Bit.
func ConvertLinearTo8Bit(v float64) uint8 {
	var scaled float64
	if v <= 0.0031308 {
		scaled = v * 12.92
	} else {
		scaled = 1.055*math.Pow(v, 1/2.4) - 0.055
	}
	return uint8(math.Round(math.Min(math.Max(scaled, 0.0), 1.0) * 255))
}

// Convert8BitToLinear converts an 8-bit sRGB encoded value to a normalised
// linear value between 0.0 and 1.0.
//
// This implementation uses an exact analytical method. If performance is
// critical, see From8Bit.
func Convert8BitToLinear(v uint8) float64 {
	vNormalised := float64(v) / 255
	if vNormalised <= 0.04045 {
		return vNormalised / 12.92
	}
	return math.Pow((vNormalised+0.055)/1.055, 2.4)
}
