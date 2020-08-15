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

// ConvertImageToRGBA is a convenience function for getting an RGBA image
// from any image. If the specified image isn’t already RGBA, a conversion is
// performed.
func ConvertImageToRGBA(img image.Image) *image.RGBA {
	switch inputImg := img.(type) {

	case *image.RGBA:
		return inputImg

	case *image.RGBA64:
		outputImg := image.NewRGBA(inputImg.Rect)

		for i := outputImg.Rect.Min.Y; i < outputImg.Rect.Max.Y; i++ {
			for j := outputImg.Rect.Min.X; j < outputImg.Rect.Max.X; j++ {
				c := inputImg.RGBA64At(j, i)
				rgba := color.RGBA{R: uint8(c.R >> 8), G: uint8(c.G >> 8), B: uint8(c.B >> 8), A: uint8(c.A >> 8)}
				outputImg.SetRGBA(j, i, rgba)
			}
		}
		return outputImg

	default:
		outputImg := image.NewRGBA(img.Bounds())
		draw.Draw(outputImg, outputImg.Rect, img, outputImg.Rect.Min, draw.Src)
		return outputImg
	}
}

// ConvertImageToRGBA64 is a convenience function for getting an RGBA64 image
// from any image. If the specified image isn’t already RGBA64, a conversion is
// performed.
func ConvertImageToRGBA64(img image.Image) *image.RGBA64 {
	switch inputImg := img.(type) {

	case *image.NRGBA:
		outputImg := image.NewRGBA64(inputImg.Rect)

		for i := outputImg.Rect.Min.Y; i < outputImg.Rect.Max.Y; i++ {
			for j := outputImg.Rect.Min.X; j < outputImg.Rect.Max.X; j++ {
				r, g, b, a := inputImg.NRGBAAt(j, i).RGBA()
				rgba64 := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
				outputImg.SetRGBA64(j, i, rgba64)
			}
		}
		return outputImg

	case *image.RGBA:
		outputImg := image.NewRGBA64(inputImg.Rect)

		for i := outputImg.Rect.Min.Y; i < outputImg.Rect.Max.Y; i++ {
			for j := outputImg.Rect.Min.X; j < outputImg.Rect.Max.X; j++ {
				r, g, b, a := inputImg.RGBAAt(j, i).RGBA()
				rgba64 := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
				outputImg.SetRGBA64(j, i, rgba64)
			}
		}
		return outputImg

	case *image.RGBA64:
		return inputImg

	case *image.YCbCr:
		outputImg := image.NewRGBA64(inputImg.Rect)

		for i := outputImg.Rect.Min.Y; i < outputImg.Rect.Max.Y; i++ {
			for j := outputImg.Rect.Min.X; j < outputImg.Rect.Max.X; j++ {
				r, g, b, _ := inputImg.YCbCrAt(j, i).RGBA()
				rgba64 := color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 65535}
				outputImg.SetRGBA64(j, i, rgba64)
			}
		}
		return outputImg

	default:
		outputImg := image.NewRGBA64(img.Bounds())
		draw.Draw(outputImg, outputImg.Rect, img, outputImg.Rect.Min, draw.Src)
		return outputImg
	}
}
