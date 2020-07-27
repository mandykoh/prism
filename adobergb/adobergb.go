package adobergb

import "math"

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

// From8Bit converts an 8-bit Adobe RGB encoded value to a normalised linear
// value between 0.0 and 1.0.
//
// This implementation uses a fast look-up table without sacrificing accuracy.
func From8Bit(srgb8 uint8) float64 {
	return adobeRGB8ToLinearLUT[srgb8]
}

// Luminance returns the perceptual luminance of the given linear RGB values.
func Luminance(r, g, b float64) float64 {
	return 0.2126*r + 0.7152*g + 0.0722*b
}

// To8Bit converts a linear value to an 8-bit Adobe RGB encoded value, clipping
// the linear value to between 0.0 and 1.0.
//
// This implementation uses a fast look-up table and is approximate. For more
// accuracy, see ConvertLinearTo8Bit.
func To8Bit(linear float64) uint8 {
	clipped := math.Min(math.Max(linear, 0), 1)
	return linearToAdobeRGB8LUT[int(math.Round(clipped*511))]
}
