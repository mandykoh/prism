package linear

import (
	"image"
	"image/color"
	"image/draw"
)

// TransformImageColor applies a colour transformation function to all pixels of
// src, writing the results to dst at its origin.
//
// src and dst may be the same image.
func TransformImageColor(dst draw.Image, src image.Image, transformColor func(color.Color) color.RGBA64) {
	bounds := src.Bounds()
	dstOffsetX := dst.Bounds().Min.X - bounds.Min.X
	dstOffsetY := dst.Bounds().Min.Y - bounds.Min.Y

	switch dstImg := dst.(type) {

	case *image.RGBA64:
		if srcImg, ok := src.(*image.RGBA64); ok {
			for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
				for j := bounds.Min.X; j < bounds.Max.X; j++ {
					dstImg.SetRGBA64(j+dstOffsetX, i+dstOffsetY, transformColor(srcImg.RGBA64At(j, i)))
				}
			}
		} else {
			for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
				for j := bounds.Min.X; j < bounds.Max.X; j++ {
					dstImg.SetRGBA64(j+dstOffsetX, i+dstOffsetY, transformColor(src.At(j, i)))
				}
			}
		}

	default:
		for i := bounds.Min.Y; i < bounds.Max.Y; i++ {
			for j := bounds.Min.X; j < bounds.Max.X; j++ {
				dst.Set(j+dstOffsetX, i+dstOffsetY, transformColor(src.At(j, i)))
			}
		}
	}
}
