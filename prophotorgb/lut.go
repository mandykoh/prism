package prophotorgb

import (
	"github.com/mandykoh/prism/linear"
	"github.com/mandykoh/prism/linear/lut"
	"sync"
)

var init8BitLUTsOnce sync.Once
var linearToEncoded8LUT []uint8
var encoded8ToLinearLUT []float32

var init16BitLUTsOnce sync.Once
var linearToEncoded16LUT []uint16
var encoded16ToLinearLUT []float32

func init() {
	init8BitLUTs()
}

// From8Bit converts an 8-bit Pro Photo RGB encoded value to a normalised linear
// value between 0.0 and 1.0.
//
// This implementation uses a fast look-up table without sacrificing accuracy.
func From8Bit(v uint8) float32 {
	return encoded8ToLinearLUT[v]
}

// From16Bit converts a 16-bit Pro Photo RGB encoded value to a normalised
// linear value between 0.0 and 1.0.
//
// This implementation uses a fast look-up table without sacrificing accuracy.
func From16Bit(v uint16) float32 {
	init16BitLUTs()
	return encoded16ToLinearLUT[v]
}

func init8BitLUTs() {
	init8BitLUTsOnce.Do(func() {
		to8BitLUT := lut.BuildLinearTo8Bit(ConvertLinearTo8Bit)
		linearToEncoded8LUT = to8BitLUT[:]

		from8BitLUT := lut.Build8BitToLinear(Convert8BitToLinear)
		encoded8ToLinearLUT = from8BitLUT[:]
	})
}

func init16BitLUTs() {
	init16BitLUTsOnce.Do(func() {
		to16BitLUT := lut.BuildLinearTo16Bit(ConvertLinearTo16Bit)
		linearToEncoded16LUT = to16BitLUT[:]

		from16BitLUT := lut.Build16BitToLinear(Convert16BitToLinear)
		encoded16ToLinearLUT = from16BitLUT[:]
	})
}

// To8Bit converts a linear value to an 8-bit Pro Photo RGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses a fast look-up table and is approximate. For more
// accuracy, see ConvertLinearTo8Bit.
func To8Bit(v float32) uint8 {
	return linearToEncoded8LUT[linear.NormalisedTo9Bit(v)]
}

// To16Bit converts a linear value to a 16-bit Pro Photo RGB encoded value,
// clipping the linear value to between 0.0 and 1.0.
//
// This implementation uses a fast look-up table and is approximate. For more
// accuracy, see ConvertLinearTo16Bit.
func To16Bit(v float32) uint16 {
	init16BitLUTs()
	return linearToEncoded16LUT[linear.NormalisedTo16Bit(v)]
}
