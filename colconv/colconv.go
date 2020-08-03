package colconv

import "image/color"

func fromPremul8To8(v, alpha uint8) uint8 {
	if alpha == 0 {
		return 0
	}
	return uint8((int(v) * 255) / int(alpha))
}

func NRGBAtoRGBA(c color.NRGBA) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
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
