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
	"runtime"
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

	nrgbaImg := image.NewNRGBA(yCbCrImg.Bounds())
	draw.Draw(nrgbaImg, nrgbaImg.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)

	rgbaImg := image.NewRGBA(yCbCrImg.Bounds())
	draw.Draw(rgbaImg, rgbaImg.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)

	rgba64Img := image.NewRGBA64(yCbCrImg.Bounds())
	draw.Draw(rgba64Img, rgba64Img.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)

	b.Run("linearisation and encoding", func(b *testing.B) {

		b.Run("from 8-bit to 16-bit and back", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				linearImg := image.NewRGBA64(rgbaImg.Rect)
				adobergb.LineariseImage(linearImg, rgbaImg, runtime.NumCPU())

				encodedImg := image.NewRGBA(linearImg.Rect)
				adobergb.EncodeImage(encodedImg, linearImg, runtime.NumCPU())
			}
		})
	})

	b.Run("between colour spaces", func(b *testing.B) {
		nrgbaOutput := image.NewNRGBA(yCbCrImg.Bounds())
		rgbaOutput := image.NewRGBA(yCbCrImg.Bounds())

		b.Run("YCbCr Adobe RGB to RGBA sRGB", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := rgbaOutput.Rect.Min.Y; i < rgbaOutput.Rect.Max.Y; i++ {
					for j := rgbaOutput.Rect.Min.X; j < rgbaOutput.Rect.Max.X; j++ {
						c := yCbCrImg.YCbCrAt(j, i)
						r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
						nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}

						ac, a := adobergb.ColorFromNRGBA(nrgba)
						sc := srgb.ColorFromXYZ(ac.ToXYZ())

						rgbaOutput.SetRGBA(j, i, sc.ToRGBA(a))
					}
				}
			}
		})

		b.Run("NRGBA Adobe RGB to RGBA sRGB", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := rgbaOutput.Rect.Min.Y; i < rgbaOutput.Rect.Max.Y; i++ {
					for j := rgbaOutput.Rect.Min.X; j < rgbaOutput.Rect.Max.X; j++ {
						c := nrgbaImg.NRGBAAt(j, i)

						ac, a := adobergb.ColorFromNRGBA(c)
						sc := srgb.ColorFromXYZ(ac.ToXYZ())

						rgbaOutput.SetRGBA(j, i, sc.ToRGBA(a))
					}
				}
			}
		})

		b.Run("RGBA Adobe RGB to RGBA sRGB", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := rgbaOutput.Rect.Min.Y; i < rgbaOutput.Rect.Max.Y; i++ {
					for j := rgbaOutput.Rect.Min.X; j < rgbaOutput.Rect.Max.X; j++ {
						c := rgbaImg.RGBAAt(j, i)

						ac, a := adobergb.ColorFromRGBA(c)
						sc := srgb.ColorFromXYZ(ac.ToXYZ())

						rgbaOutput.SetRGBA(j, i, sc.ToRGBA(a))
					}
				}
			}
		})

		b.Run("YCbCr Adobe RGB to NRGBA sRGB", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := nrgbaOutput.Rect.Min.Y; i < nrgbaOutput.Rect.Max.Y; i++ {
					for j := nrgbaOutput.Rect.Min.X; j < nrgbaOutput.Rect.Max.X; j++ {
						c := yCbCrImg.YCbCrAt(j, i)
						r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
						nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}

						ac, a := adobergb.ColorFromNRGBA(nrgba)
						sc := srgb.ColorFromXYZ(ac.ToXYZ())

						nrgbaOutput.SetNRGBA(j, i, sc.ToNRGBA(a))
					}
				}
			}
		})

		b.Run("NRGBA Adobe RGB to NRGBA sRGB", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := nrgbaOutput.Rect.Min.Y; i < nrgbaOutput.Rect.Max.Y; i++ {
					for j := nrgbaOutput.Rect.Min.X; j < nrgbaOutput.Rect.Max.X; j++ {
						c := nrgbaImg.NRGBAAt(j, i)

						ac, a := adobergb.ColorFromNRGBA(c)
						sc := srgb.ColorFromXYZ(ac.ToXYZ())

						nrgbaOutput.SetNRGBA(j, i, sc.ToNRGBA(a))
					}
				}
			}
		})

		b.Run("RGBA Adobe RGB to NRGBA sRGB", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := nrgbaOutput.Rect.Min.Y; i < nrgbaOutput.Rect.Max.Y; i++ {
					for j := nrgbaOutput.Rect.Min.X; j < nrgbaOutput.Rect.Max.X; j++ {
						c := rgbaImg.RGBAAt(j, i)

						ac, a := adobergb.ColorFromRGBA(c)
						sc := srgb.ColorFromXYZ(ac.ToXYZ())

						nrgbaOutput.SetNRGBA(j, i, sc.ToNRGBA(a))
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
				for i := nrgbaOutput.Rect.Min.Y; i < nrgbaOutput.Rect.Max.Y; i++ {
					for j := nrgbaOutput.Rect.Min.X; j < nrgbaOutput.Rect.Max.X; j++ {
						c := yCbCrImg.YCbCrAt(j, i)
						r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
						nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}

						ac, a := adobergb.ColorFromNRGBA(nrgba)
						pc := prophotorgb.ColorFromXYZ(adaptation.Apply(ac.ToXYZ()))

						nrgbaOutput.SetNRGBA(j, i, pc.ToNRGBA(a))
					}
				}
			}
		})

		b.Run("NRGBA Adobe RGB to NRGBA Pro Photo", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := nrgbaOutput.Rect.Min.Y; i < nrgbaOutput.Rect.Max.Y; i++ {
					for j := nrgbaOutput.Rect.Min.X; j < nrgbaOutput.Rect.Max.X; j++ {
						c := nrgbaImg.NRGBAAt(j, i)

						ac, a := adobergb.ColorFromNRGBA(c)
						pc := prophotorgb.ColorFromXYZ(adaptation.Apply(ac.ToXYZ()))

						nrgbaOutput.SetNRGBA(j, i, pc.ToNRGBA(a))
					}
				}
			}
		})

		b.Run("RGBA Adobe RGB to NRGBA Pro Photo", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				for i := nrgbaOutput.Rect.Min.Y; i < nrgbaOutput.Rect.Max.Y; i++ {
					for j := nrgbaOutput.Rect.Min.X; j < nrgbaOutput.Rect.Max.X; j++ {
						c := rgbaImg.RGBAAt(j, i)

						ac, a := adobergb.ColorFromRGBA(c)
						pc := prophotorgb.ColorFromXYZ(adaptation.Apply(ac.ToXYZ()))

						nrgbaOutput.SetNRGBA(j, i, pc.ToNRGBA(a))
					}
				}
			}
		})
	})

	b.Run("between colour models", func(b *testing.B) {

		b.Run("YCbCr to NRGBA non-colour managed draw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				output := image.NewNRGBA(yCbCrImg.Rect)
				draw.Draw(output, output.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)
			}
		})

		b.Run("YCbCr to RGBA non-colour managed draw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				output := image.NewRGBA(yCbCrImg.Rect)
				draw.Draw(output, output.Rect, yCbCrImg, yCbCrImg.Rect.Min, draw.Src)
			}
		})

		b.Run("NRGBA to RGBA non-colour managed draw", func(b *testing.B) {
			// NRGBA to RGBA is more correctly done via a colour space as tonal
			// response encoding is applied on top of alpha premultiplication.
			for i := 0; i < b.N; i++ {
				output := image.NewRGBA(nrgbaImg.Rect)
				draw.Draw(output, output.Rect, nrgbaImg, nrgbaImg.Rect.Min, draw.Src)
			}
		})

		b.Run("RGBA to NRGBA non-colour managed draw", func(b *testing.B) {
			// RGBA to NRGBA is more correctly done via a colour space as tonal
			// response encoding is applied on top of alpha premultiplication.
			for i := 0; i < b.N; i++ {
				output := image.NewNRGBA(rgbaImg.Rect)
				draw.Draw(output, output.Rect, rgbaImg, rgbaImg.Rect.Min, draw.Src)
			}
		})

		b.Run("YCbCr to NRGBA non-colour managed pixel copy", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				output := image.NewNRGBA(yCbCrImg.Rect)

				for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
					for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
						c := yCbCrImg.YCbCrAt(j, i)
						r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
						nrgba := color.NRGBA{R: r, G: g, B: b, A: 255}
						output.SetNRGBA(j, i, nrgba)
					}
				}
			}
		})

		b.Run("YCbCr to RGBA non-colour managed pixel copy", func(b *testing.B) {
			for iteration := 0; iteration < b.N; iteration++ {
				output := image.NewRGBA(yCbCrImg.Rect)

				for i := output.Rect.Min.Y; i < output.Rect.Max.Y; i++ {
					for j := output.Rect.Min.X; j < output.Rect.Max.X; j++ {
						c := yCbCrImg.YCbCrAt(j, i)
						r, g, b := color.YCbCrToRGB(c.Y, c.Cb, c.Cr)
						rgba := color.RGBA{R: r, G: g, B: b, A: 255}
						output.SetRGBA(j, i, rgba)
					}
				}
			}
		})

		b.Run("NRGBA to NRGBA with ConvertImageToNRGBA()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToNRGBA(nrgbaImg, runtime.NumCPU())
			}
		})

		b.Run("RGBA to NRGBA with ConvertImageToNRGBA()", func(b *testing.B) {
			// RGBA to NRGBA is more correctly done via a colour space as tonal
			// response encoding is applied on top of alpha premultiplication.
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToNRGBA(rgbaImg, runtime.NumCPU())
			}
		})

		b.Run("YCbCr to NRGBA with ConvertImageToNRGBA()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToNRGBA(yCbCrImg, runtime.NumCPU())
			}
		})

		b.Run("NRGBA to RGBA with ConvertImageToRGBA()", func(b *testing.B) {
			// RGBA to NRGBA is more correctly done via a colour space as tonal
			// response encoding is applied on top of alpha premultiplication.
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA(nrgbaImg, runtime.NumCPU())
			}
		})

		b.Run("RGBA64 to RGBA with ConvertImageToRGBA()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA(rgba64Img, runtime.NumCPU())
			}
		})

		b.Run("RGBA to RGBA with ConvertImageToRGBA()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA(rgbaImg, runtime.NumCPU())
			}
		})

		b.Run("YCbCr to RGBA with ConvertImageToRGBA()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA(yCbCrImg, runtime.NumCPU())
			}
		})

		b.Run("NRGBA to RGBA64 with ConvertImageToRGBA64()", func(b *testing.B) {
			// RGBA to NRGBA is more correctly done via a colour space as tonal
			// response encoding is applied on top of alpha premultiplication.
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA64(nrgbaImg, runtime.NumCPU())
			}
		})

		b.Run("RGBA64 to RGBA64 with ConvertImageToRGBA64()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA64(rgba64Img, runtime.NumCPU())
			}
		})

		b.Run("RGBA to RGBA64 with ConvertImageToRGBA64()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA64(rgbaImg, runtime.NumCPU())
			}
		})

		b.Run("YCbCr to RGBA64 with ConvertImageToRGBA64()", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ConvertImageToRGBA64(yCbCrImg, runtime.NumCPU())
			}
		})
	})
}
