package autometa_test

import (
	"fmt"
	"github.com/mandykoh/prism/meta"
	"github.com/mandykoh/prism/meta/autometa"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func printMetadata(md *meta.Data, img image.Image) {
	fmt.Printf("Format: %s\n", md.Format)
	fmt.Printf("BitsPerComponent: %d\n", md.BitsPerComponent)
	fmt.Printf("PixelHeight: %d\n", md.PixelHeight)
	fmt.Printf("PixelWidth: %d\n", md.PixelWidth)

	fmt.Printf("Actual image height: %d\n", img.Bounds().Dy())
	fmt.Printf("Actual image width: %d\n", img.Bounds().Dx())
}

func ExampleLoad_basicJPEGMetadata() {
	inFile, err := os.Open("../../test-images/pizza-rgb8-srgb.jpg")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	md, imgStream, err := autometa.Load(inFile)
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(imgStream)
	if err != nil {
		panic(err)
	}

	printMetadata(md, img)

	// Output:
	// Format: JPEG
	// BitsPerComponent: 8
	// PixelHeight: 1200
	// PixelWidth: 1200
	// Actual image height: 1200
	// Actual image width: 1200
}

func ExampleLoad_basicPNGMetadata() {
	inFile, err := os.Open("../../test-images/pizza-rgb8-srgb.png")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	md, imgStream, err := autometa.Load(inFile)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(imgStream)
	if err != nil {
		panic(err)
	}

	printMetadata(md, img)

	// Output:
	// Format: PNG
	// BitsPerComponent: 8
	// PixelHeight: 1200
	// PixelWidth: 1200
	// Actual image height: 1200
	// Actual image width: 1200
}
