// Package prism provides a set of tools for colour management and conversion.
// Subpackages provide support for encoding/decoding image pixel data in
// specific colour spaces, and conversions between those spaces.
package prism

import (
	"image"
	"image/color"
	"image/draw"
)

// ConvertImageToNRGBA is a convenience function for getting an NRGBA image from
// any image. If the specified image isn’t already NRGBA, a conversion is
// performed.
func ConvertImageToNRGBA(img image.Image) *image.NRGBA {
	switch inputImg := img.(type) {

	case *image.NRGBA:
		return inputImg

	case *image.YCbCr:
		outputImg := image.NewNRGBA(inputImg.Rect)

		for i := outputImg.Rect.Min.Y; i < outputImg.Rect.Max.Y; i++ {
			for j := outputImg.Rect.Min.X; j < outputImg.Rect.Max.X; j++ {
				c := inputImg.YCbCrAt(j, i)
				r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
				nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}
				outputImg.SetNRGBA(j, i, nrgba)
			}
		}
		return outputImg

	default:
		outputImg := image.NewNRGBA(img.Bounds())
		draw.Draw(outputImg, outputImg.Rect, img, outputImg.Rect.Min, draw.Src)
		return outputImg
	}
}
