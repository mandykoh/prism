package adobergb

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

	yCbCrImg := loadImage("../test-images/pizza-rgb8-adobergb.jpg")

	nrgbaImg := image.NewNRGBA(yCbCrImg.Bounds())
	LineariseImage(nrgbaImg, yCbCrImg)

	rgbaImg := image.NewRGBA(yCbCrImg.Bounds())
	LineariseImage(rgbaImg, yCbCrImg)

	rgba64Img := image.NewRGBA64(yCbCrImg.Bounds())
	LineariseImage(rgba64Img, yCbCrImg)

	b.Run("with NRGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewNRGBA(nrgbaImg.Rect)
			EncodeImage(result, nrgbaImg)
		}
	})

	b.Run("with RGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA(rgbaImg.Rect)
			EncodeImage(result, rgbaImg)
		}
	})

	b.Run("with RGBA64 image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA64(rgba64Img.Rect)
			EncodeImage(result, rgba64Img)
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

	yCbCrImg := loadImage("../test-images/pizza-rgb8-adobergb.jpg")
	nrgbaImg := prism.ConvertImageToNRGBA(yCbCrImg)
	rgbaImg := prism.ConvertImageToRGBA(yCbCrImg)
	rgba64Img := prism.ConvertImageToRGBA64(yCbCrImg)

	b.Run("with NRGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewNRGBA(nrgbaImg.Rect)
			LineariseImage(result, nrgbaImg)
		}
	})

	b.Run("with RGBA image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA(rgbaImg.Rect)
			LineariseImage(result, rgbaImg)
		}
	})

	b.Run("with RGBA64 image", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := image.NewRGBA64(rgba64Img.Rect)
			LineariseImage(result, rgba64Img)
		}
	})
}
