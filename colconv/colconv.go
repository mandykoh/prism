package colconv

import "image/color"

func from8ToPremul8(v, alpha uint8) uint8 {
	return uint8(int(v) * int(alpha) / 255)
}

func fromPremul8To8(v, alpha uint8) uint8 {
	if alpha == 0 {
		return 0
	}
	return uint8((int(v) * 255) / int(alpha))
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
	return color.NRGBA{
		R: fromPremul8To8(c.R, c.A),
		G: fromPremul8To8(c.G, c.A),
		B: fromPremul8To8(c.B, c.A),
		A: c.A,
	}
}
