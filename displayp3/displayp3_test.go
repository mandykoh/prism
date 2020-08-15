package displayp3

import (
	"github.com/mandykoh/prism"
	"image"
	_ "image/jpeg"
	"os"
	"testing"
)

func BenchmarkLineariseImage(b *testing.B) {

	loadImage := func(path string) image.Image {
		inFile, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer inFile.Close()

		img, _, err := image.Decode(inFile)
		if err != nil {
			panic(err)
		}

		return img
	}

	yCbCrImg := loadImage("../test-images/pizza-rgb8-displayp3.jpg")
	nrgbaImg := prism.ConvertImageToNRGBA(yCbCrImg)
	rgbaImg := prism.ConvertImageToRGBA(yCbCrImg)
	rgba64Img := prism.ConvertImageToRGBA64(yCbCrImg)

	b.Run("with NRGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LineariseImage(nrgbaImg)
		}
	})

	b.Run("with RGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LineariseImage(rgbaImg)
		}
	})

	b.Run("with RGBA64 image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LineariseImage(rgba64Img)
		}
	})
}
