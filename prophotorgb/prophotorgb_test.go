package prophotorgb

import (
	"github.com/mandykoh/prism"
	"image"
	_ "image/jpeg"
	"os"
	"testing"
)

func BenchmarkEncodeImage(b *testing.B) {

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

	yCbCrImg := loadImage("../test-images/pizza-rgb8-prophotorgb.jpg")

	nrgbaImg := prism.ConvertImageToNRGBA(yCbCrImg)
	LineariseImage(nrgbaImg)

	rgbaImg := prism.ConvertImageToRGBA(yCbCrImg)
	LineariseImage(rgbaImg)

	rgba64Img := prism.ConvertImageToRGBA64(yCbCrImg)
	LineariseImage(rgba64Img)

	b.Run("with NRGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EncodeImage(nrgbaImg)
		}
	})

	b.Run("with RGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EncodeImage(rgbaImg)
		}
	})

	b.Run("with RGBA64 image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EncodeImage(rgba64Img)
		}
	})
}

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

	yCbCrImg := loadImage("../test-images/pizza-rgb8-prophotorgb.jpg")
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

func TestConvertLinearTo8Bit(t *testing.T) {

	t.Run("clips linear values to between 0 and 1", func(t *testing.T) {
		if expected, actual := ConvertLinearTo8Bit(0), ConvertLinearTo8Bit(-0.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
		if expected, actual := ConvertLinearTo8Bit(1), ConvertLinearTo8Bit(1.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
	})
}
