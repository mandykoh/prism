package prophotorgb

import (
	"github.com/mandykoh/prism"
	"image"
	_ "image/jpeg"
	"os"
	"runtime"
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

	nrgbaImg := image.NewNRGBA(yCbCrImg.Bounds())
	LineariseImage(nrgbaImg, yCbCrImg, runtime.NumCPU())

	rgbaImg := image.NewRGBA(yCbCrImg.Bounds())
	LineariseImage(rgbaImg, yCbCrImg, runtime.NumCPU())

	rgba64Img := image.NewRGBA64(yCbCrImg.Bounds())
	LineariseImage(rgba64Img, yCbCrImg, runtime.NumCPU())

	b.Run("with NRGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewNRGBA(nrgbaImg.Rect)
			EncodeImage(result, nrgbaImg, runtime.NumCPU())
		}
	})

	b.Run("with RGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA(rgbaImg.Rect)
			EncodeImage(result, rgbaImg, runtime.NumCPU())
		}
	})

	b.Run("with RGBA64 image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA64(rgba64Img.Rect)
			EncodeImage(result, rgba64Img, runtime.NumCPU())
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
	nrgbaImg := prism.ConvertImageToNRGBA(yCbCrImg, runtime.NumCPU())
	rgbaImg := prism.ConvertImageToRGBA(yCbCrImg, runtime.NumCPU())
	rgba64Img := prism.ConvertImageToRGBA64(yCbCrImg, runtime.NumCPU())

	b.Run("with NRGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewNRGBA(nrgbaImg.Rect)
			LineariseImage(result, nrgbaImg, runtime.NumCPU())
		}
	})

	b.Run("with RGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA(rgbaImg.Rect)
			LineariseImage(result, rgbaImg, runtime.NumCPU())
		}
	})

	b.Run("with RGBA64 image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA64(rgba64Img.Rect)
			LineariseImage(result, rgba64Img, runtime.NumCPU())
		}
	})
}
