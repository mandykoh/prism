package srgb

import (
	"github.com/mandykoh/prism/ciexyy"
	"math"
)

var PrimaryRed = ciexyy.Color{X: 0.64, Y: 0.33, YY: 1}
var PrimaryGreen = ciexyy.Color{X: 0.3, Y: 0.6, YY: 1}
var PrimaryBlue = ciexyy.Color{X: 0.15, Y: 0.06, YY: 1}
var StandardWhitePoint = ciexyy.D65

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
	if vNormalised <= 0.0031308*12.92 {
		return vNormalised / 12.92
	}
	return math.Pow((vNormalised+0.055)/1.055, 2.4)
}
