package autometa_test

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/mandykoh/prism/meta"
	"github.com/mandykoh/prism/meta/autometa"
)

func printMetadata(md *meta.Data, img image.Image) {
	profile, err := md.ICCProfile()
	if err != nil {
		panic(err)
	}

	var profileDescription = "sRGB IEC61966-2.1"
	if profile != nil {
		profileDescription, err = profile.Description()
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Format: %s\n", md.Format)
	fmt.Printf("Profile: %s\n", profileDescription)
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
	// Profile: sRGB IEC61966-2.1
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
	// Profile: sRGB IEC61966-2.1
	// BitsPerComponent: 8
	// PixelHeight: 1200
	// PixelWidth: 1200
	// Actual image height: 1200
	// Actual image width: 1200
}

func ExampleLoad_iOSExportJPEGMetadata() {
	inFile, err := os.Open("../../test-images/ios-export.jpg")
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
	// Profile: Display P3
	// BitsPerComponent: 8
	// PixelHeight: 400
	// PixelWidth: 600
	// Actual image height: 400
	// Actual image width: 600
}
