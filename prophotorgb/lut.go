package prophotorgb

import (
	"math"
	"sync"
)

var init8BitLUTsOnce sync.Once
var linearToEncoded8LUT []uint8
var encoded8ToLinearLUT []float32

var init16BitLUTsOnce sync.Once
var linearToEncoded16LUT []uint16

// From8Bit converts an 8-bit Pro Photo RGB encoded value to a normalised linear
// value between 0.0 and 1.0.
//
// This implementation uses a fast look-up table without sacrificing accuracy.
func From8Bit(v uint8) float32 {
	init8BitLUTs()
	return encoded8ToLinearLUT[v]
}

func init8BitLUTs() {
	init8BitLUTsOnce.Do(func() {
		to8BitLUT := make([]uint8, 512)
		for i := range to8BitLUT {
			to8BitLUT[i] = ConvertLinearTo8Bit(float32(i) / 511)
		}
		linearToEncoded8LUT = to8BitLUT

		from8BitLUT := make([]float32, 256)
		for i := range from8BitLUT {
			from8BitLUT[i] = Convert8BitToLinear(uint8(i))
		}
		encoded8ToLinearLUT = from8BitLUT
	})
}

func init16BitLUTs() {
	init16BitLUTsOnce.Do(func() {
		to16BitLUT := make([]uint16, 65536)
		for i := range to16BitLUT {
			to16BitLUT[i] = ConvertLinearTo16Bit(float32(i) / 65535)
		}
		linearToEncoded16LUT = to16BitLUT
	})
}

// To8Bit converts a linear value to an 8-bit Pro Photo RGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses a fast look-up table and is approximate. For more
// accuracy, see ConvertLinearTo8Bit.
func To8Bit(linear float32) uint8 {
	init8BitLUTs()
	clipped := math.Min(math.Max(float64(linear), 0), 1)
	return linearToEncoded8LUT[int(math.Round(clipped*511))]
}

// To16Bit converts a linear value to a 16-bit Pro Photo RGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses a fast look-up table and is approximate. For more
// accuracy, see ConvertLinearTo16Bit.
func To16Bit(linear float32) uint16 {
	init16BitLUTs()
	clipped := math.Min(math.Max(float64(linear), 0), 1)
	return linearToEncoded16LUT[int(math.Round(clipped*65535))]
}
