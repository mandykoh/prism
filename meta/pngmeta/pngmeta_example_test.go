package pngmeta_test

import (
	"fmt"
	"github.com/mandykoh/prism/meta"
	"github.com/mandykoh/prism/meta/pngmeta"
	"image"
	"image/png"
	"os"
)

func printICCProfile(md *meta.Data) {
	profile, err := md.ICCProfile()
	if err != nil {
		panic(err)
	}

	fmt.Printf("ProfileSize: %d\n", profile.Header.ProfileSize)
	fmt.Printf("PreferredCMM: %v\n", profile.Header.PreferredCMM)
	fmt.Printf("ProfileVersion: %s\n", profile.Header.Version)
	fmt.Printf("DeviceClass: %s\n", profile.Header.DeviceClass)
	fmt.Printf("DataColorSpace: %s\n", profile.Header.DataColorSpace)
	fmt.Printf("PCS: %s\n", profile.Header.ProfileConnectionSpace)
	fmt.Printf("CreatedAt: %v\n", profile.Header.CreatedAt)
	fmt.Printf("PrimaryPlatform: %v\n", profile.Header.PrimaryPlatform)
	fmt.Printf("Embedded: %v\n", profile.Header.Embedded)
	fmt.Printf("DependsOnEmbeddedData: %v\n", profile.Header.DependsOnEmbeddedData)
	fmt.Printf("DeviceManufacturer: %s\n", profile.Header.DeviceManufacturer)
	fmt.Printf("DeviceModel: %s\n", profile.Header.DeviceModel)
	fmt.Printf("DeviceAttributes: %064b\n", profile.Header.DeviceAttributes)
	fmt.Printf("RenderingIntent: %v\n", profile.Header.RenderingIntent)
	fmt.Printf("PCSIlluminant: %v\n", profile.Header.PCSIlluminant)
	fmt.Printf("ProfileCreator: %v\n", profile.Header.ProfileCreator)
	fmt.Printf("ProfileID: %0x\n", profile.Header.ProfileID)

	if desc, err := profile.Description(); err != nil {
		panic(err)
	} else {
		fmt.Printf("Description: %s\n", desc)
	}
}

func printMetadata(md *meta.Data, img image.Image) {
	fmt.Printf("BitsPerComponent: %d\n", md.BitsPerComponent)
	fmt.Printf("PixelHeight: %d\n", md.PixelHeight)
	fmt.Printf("PixelWidth: %d\n", md.PixelWidth)

	fmt.Printf("Actual image height: %d\n", img.Bounds().Dy())
	fmt.Printf("Actual image width: %d\n", img.Bounds().Dx())
}

func ExampleLoad_basicPNGMetadata() {
	inFile, err := os.Open("../../test-images/pizza-rgb8-srgb.png")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	md, imgStream, err := pngmeta.Load(inFile)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(imgStream)
	if err != nil {
		panic(err)
	}

	printMetadata(md, img)

	// Output:
	// BitsPerComponent: 8
	// PixelHeight: 1200
	// PixelWidth: 1200
	// Actual image height: 1200
	// Actual image width: 1200
}

func ExampleLoad_embeddedICCv4() {
	inFile, err := os.Open("../../test-images/pizza-rgb8-srgb.png")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	md, imgStream, err := pngmeta.Load(inFile)
	if err != nil {
		panic(err)
	}

	_, err = png.Decode(imgStream)
	if err != nil {
		panic(err)
	}

	printICCProfile(md)

	// Output:
	// ProfileSize: 596
	// PreferredCMM: 'lcms'
	// ProfileVersion: 4.3.0
	// DeviceClass: Display
	// DataColorSpace: RGB
	// PCS: XYZ
	// CreatedAt: 2020-08-12 21:45:21 +0000 UTC
	// PrimaryPlatform: Apple Computer, Inc.
	// Embedded: false
	// DependsOnEmbeddedData: false
	// DeviceManufacturer: '    '
	// DeviceModel: '    '
	// DeviceAttributes: 0000000000000000000000000000000000000000000000000000000000000000
	// RenderingIntent: Perceptual
	// PCSIlluminant: [63190 65536 54061]
	// ProfileCreator: 'lcms'
	// ProfileID: 00000000000000000000000000000000
	// Description: sRGB IEC61966-2.1
}
