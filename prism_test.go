package prism

import (
	"github.com/mandykoh/prism/adobergb"
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/prophotorgb"
	"github.com/mandykoh/prism/srgb"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func loadImage(path string) image.Image {
	imgFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(imgFile)
	_ = imgFile.Close()
	if err != nil {
		panic(err)
	}

	return img
}

func BenchmarkColorConversion(b *testing.B) {
	yCbCrImg := loadImage("test-images/pizza-rgb8-adobergb.jpg").(*image.YCbCr)

	output := image.NewNRGBA(yCbCrImg.Bounds())

	b.Run("YCbCr Adobe RGB to NRGBA sRGB", func(b *testing.B) {
		for iteration := 0; iteration < b.N; iteration++ {
			for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
				for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
					c := yCbCrImg.YCbCrAt(j, i)
					r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
					nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}

					ac, a := adobergb.ColorFromNRGBA(nrgba)
					sc := srgb.ColorFromXYZ(ac.ToXYZ())

					output.SetNRGBA(j, i, sc.ToNRGBA(a))
				}
			}
		}
	})

	nrgbaImg := image.NewNRGBA(yCbCrImg.Bounds())
	draw.Draw(nrgbaImg, nrgbaImg.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)

	b.Run("NRGBA Adobe RGB to NRGBA sRGB", func(b *testing.B) {
		for iteration := 0; iteration < b.N; iteration++ {
			for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
				for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
					c := nrgbaImg.NRGBAAt(j, i)

					ac, a := adobergb.ColorFromNRGBA(c)
					sc := srgb.ColorFromXYZ(ac.ToXYZ())

					output.SetNRGBA(j, i, sc.ToNRGBA(a))
				}
			}
		}
	})

	rgbaImg := image.NewRGBA(yCbCrImg.Bounds())
	draw.Draw(rgbaImg, rgbaImg.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)

	b.Run("RGBA Adobe RGB to NRGBA sRGB", func(b *testing.B) {
		for iteration := 0; iteration < b.N; iteration++ {
			for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
				for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
					c := rgbaImg.RGBAAt(j, i)

					ac, a := adobergb.ColorFromRGBA(c)
					sc := srgb.ColorFromXYZ(ac.ToXYZ())

					output.SetNRGBA(j, i, sc.ToNRGBA(a))
				}
			}
		}
	})

	adaptation := ciexyz.AdaptBetweenXYYWhitePoints(
		adobergb.StandardWhitePoint,
		prophotorgb.StandardWhitePoint,
	)

	b.Run("YCbCr Adobe RGB to NRGBA Pro Photo", func(b *testing.B) {
		for iteration := 0; iteration < b.N; iteration++ {
			for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
				for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
					c := yCbCrImg.YCbCrAt(j, i)
					r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
					nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}

					ac, a := adobergb.ColorFromNRGBA(nrgba)
					pc := prophotorgb.ColorFromXYZ(adaptation.Apply(ac.ToXYZ()))

					output.SetNRGBA(j, i, pc.ToNRGBA(a))
				}
			}
		}
	})

	b.Run("NRGBA Adobe RGB to NRGBA Pro Photo", func(b *testing.B) {
		for iteration := 0; iteration < b.N; iteration++ {
			for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
				for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
					c := nrgbaImg.NRGBAAt(j, i)

					ac, a := adobergb.ColorFromNRGBA(c)
					pc := prophotorgb.ColorFromXYZ(adaptation.Apply(ac.ToXYZ()))

					output.SetNRGBA(j, i, pc.ToNRGBA(a))
				}
			}
		}
	})

	b.Run("RGBA Adobe RGB to NRGBA Pro Photo", func(b *testing.B) {
		for iteration := 0; iteration < b.N; iteration++ {
			for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
				for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
					c := rgbaImg.RGBAAt(j, i)

					ac, a := adobergb.ColorFromRGBA(c)
					pc := prophotorgb.ColorFromXYZ(adaptation.Apply(ac.ToXYZ()))

					output.SetNRGBA(j, i, pc.ToNRGBA(a))
				}
			}
		}
	})
}
