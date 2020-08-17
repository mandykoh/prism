package lut

func BuildLinearTo8Bit(conversion func(float32) uint8) [512]uint8 {
	to8BitLUT := [512]uint8{}
	for i := range to8BitLUT {
		to8BitLUT[i] = conversion(float32(i) / 511)
	}
	return to8BitLUT
}

func BuildLinearTo16Bit(conversion func(float32) uint16) [65536]uint16 {
	to16BitLUT := [65536]uint16{}
	for i := range to16BitLUT {
		to16BitLUT[i] = conversion(float32(i) / 65535)
	}
	return to16BitLUT
}

func Build8BitToLinear(conversion func(uint8) float32) [256]float32 {
	from8BitLUT := [256]float32{}
	for i := range from8BitLUT {
		from8BitLUT[i] = conversion(uint8(i))
	}
	return from8BitLUT
}

func Build16BitToLinear(conversion func(uint16) float32) [65536]float32 {
	from16BitLUT := [65536]float32{}
	for i := range from16BitLUT {
		from16BitLUT[i] = conversion(uint16(i))
	}
	return from16BitLUT
}
