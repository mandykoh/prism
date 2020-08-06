package colconv

import "image/color"

func from8ToPremul8(v, alpha uint8) uint8 {
	return uint8(uint32(v) * uint32(alpha) / 255)
}

func NRGBAtoRGBA(c color.NRGBA) color.RGBA {
	return color.RGBA{
		R: from8ToPremul8(c.R, c.A),
		G: from8ToPremul8(c.G, c.A),
		B: from8ToPremul8(c.B, c.A),
		A: c.A,
	}
}

func RGBAtoNRGBA(c color.RGBA) color.NRGBA {
	if c.A == 0 {
		return color.NRGBA{}
	}

	return color.NRGBA{
		R: uint8((uint32(c.R) * 255) / uint32(c.A)),
		G: uint8((uint32(c.G) * 255) / uint32(c.A)),
		B: uint8((uint32(c.B) * 255) / uint32(c.A)),
		A: c.A,
	}
}
