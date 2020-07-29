package prism_test

import (
	"fmt"
	"github.com/mandykoh/prism/adobergb"
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/prophotorgb"
	"github.com/mandykoh/prism/srgb"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func loadImage(path string) *image.NRGBA {
	imgFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	nrgbaImg := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))
	draw.Draw(nrgbaImg, nrgbaImg.Rect, img, img.Bounds().Min, draw.Src)

	return nrgbaImg
}

func compare(img1, img2 *image.NRGBA, threshold int) float64 {
	diffCount := 0

	for i := img1.Rect.Min.Y; i < img1.Rect.Max.Y; i++ {
		for j := img1.Rect.Min.X; j < img1.Rect.Max.X; j++ {
			c1 := img1.NRGBAAt(j, i)
			d1 := [4]int{int(c1.R), int(c1.G), int(c1.B), int(c1.A)}

			c2 := img2.NRGBAAt(j, i)
			d2 := [4]int{int(c2.R), int(c2.G), int(c2.B), int(c2.A)}

			diff := 0
			for k := range d1 {
				if d1[k] > d2[k] {
					diff += d1[k] - d2[k]
				} else {
					diff += d2[k] - d1[k]
				}
			}

			if diff > threshold {
				diffCount++
			}
		}
	}

	return float64(diffCount) / float64(img1.Rect.Dx()*img1.Rect.Dy())
}

func writeImage(path string, img image.Image) {
	_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)

	imgFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	err = png.Encode(imgFile, img)
	if err != nil {
		panic(err)
	}

	log.Printf("Output written to %s", path)
}

func Example_ConvertAdobeRGBToSRGB() {
	referenceImg := loadImage("test-images/pizza-rgb8-srgb.jpg")
	inputImg := loadImage("test-images/pizza-rgb8-adobergb.jpg")

	convertedImg := image.NewNRGBA(inputImg.Rect)
	for i := inputImg.Rect.Min.Y; i < inputImg.Rect.Max.Y; i++ {
		for j := inputImg.Rect.Min.X; j < inputImg.Rect.Max.X; j++ {
			inCol, alpha := adobergb.ColorFromNRGBA(inputImg.NRGBAAt(j, i))
			outCol := srgb.ColorFromXYZ(inCol.ToXYZ())
			convertedImg.SetNRGBA(j, i, outCol.ToNRGBA(alpha))
		}
	}

	writeImage("example-output/adobergb-to-srgb.png", convertedImg)

	if difference := compare(convertedImg, referenceImg, 5); difference > 0.01 {
		fmt.Printf("Images differ by %.2f%% of pixels exceeding difference threshold", difference*100)
	} else {
		fmt.Printf("Images match")
	}

	// Output: Images match
}

func Example_ConvertProPhotoRGBToSRGB() {
	referenceImg := loadImage("test-images/pizza-rgb8-srgb.jpg")
	inputImg := loadImage("test-images/pizza-rgb8-prophotorgb.jpg")

	adaptation := ciexyz.AdaptBetweenWhitePoints(
		prophotorgb.StandardWhitePoint,
		srgb.StandardWhitePoint,
	)

	convertedImg := image.NewNRGBA(inputImg.Rect)
	for i := inputImg.Rect.Min.Y; i < inputImg.Rect.Max.Y; i++ {
		for j := inputImg.Rect.Min.X; j < inputImg.Rect.Max.X; j++ {
			inCol, alpha := prophotorgb.ColorFromNRGBA(inputImg.NRGBAAt(j, i))

			xyz := inCol.ToXYZ()
			xyz = adaptation.Apply(xyz)

			outCol := srgb.ColorFromXYZ(xyz)
			convertedImg.SetNRGBA(j, i, outCol.ToNRGBA(alpha))
		}
	}

	writeImage("example-output/prophotorgb-to-srgb.png", convertedImg)

	if difference := compare(convertedImg, referenceImg, 5); difference > 0.015 {
		fmt.Printf("Images differ by %.2f%% of pixels exceeding difference threshold", difference*100)
	} else {
		fmt.Printf("Images match")
	}

	// Output: Images match
}

func Example_ConvertSRGBToAdobeRGB() {
	referenceImg := loadImage("test-images/pizza-rgb8-adobergb.jpg")
	inputImg := loadImage("test-images/pizza-rgb8-srgb.jpg")

	convertedImg := image.NewNRGBA(inputImg.Rect)
	for i := inputImg.Rect.Min.Y; i < inputImg.Rect.Max.Y; i++ {
		for j := inputImg.Rect.Min.X; j < inputImg.Rect.Max.X; j++ {
			inCol, alpha := srgb.ColorFromNRGBA(inputImg.NRGBAAt(j, i))
			outCol := adobergb.ColorFromXYZ(inCol.ToXYZ())
			convertedImg.SetNRGBA(j, i, outCol.ToNRGBA(alpha))
		}
	}

	// Output will be written without an embedded colour profile (software used
	// to examine this image will assume sRGB unless told otherwise).
	//writeImage("example-output/srgb-to-adobergb.png", convertedImg)

	if difference := compare(convertedImg, referenceImg, 4); difference > 0.01 {
		fmt.Printf("Images differ by %.2f%% of pixels exceeding difference threshold", difference*100)
	} else {
		fmt.Printf("Images match")
	}

	// Output: Images match
}
