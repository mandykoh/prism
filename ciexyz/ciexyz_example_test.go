package ciexyz_test

import (
	"fmt"
	"github.com/mandykoh/prism/adobergb"
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/displayp3"
	"github.com/mandykoh/prism/prophotorgb"
	"github.com/mandykoh/prism/srgb"
)

func ExampleTransformFromXYZForXYYPrimaries_generateAdobeRGBMatrix() {
	transform := ciexyz.TransformFromXYZForXYYPrimaries(
		adobergb.PrimaryRed,
		adobergb.PrimaryGreen,
		adobergb.PrimaryBlue,
		adobergb.StandardWhitePoint)

	fmt.Printf("R = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*2.041591 + c.Y*-0.565008 + c.Z*-0.344732
	// G = c.X*-0.969224 + c.Y*1.875930 + c.Z*0.041554
	// B = c.X*0.013446 + c.Y*-0.118381 + c.Z*1.015338
}

func ExampleTransformToXYZForXYYPrimaries_generateAdobeRGBMatrix() {
	transform := ciexyz.TransformToXYZForXYYPrimaries(
		adobergb.PrimaryRed,
		adobergb.PrimaryGreen,
		adobergb.PrimaryBlue,
		adobergb.StandardWhitePoint)

	fmt.Printf("X = c.R*%f + c.G*%f + c.B*%f\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("Y = c.R*%f + c.G*%f + c.B*%f\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("Z = c.R*%f + c.G*%f + c.B*%f\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// X = c.R*0.576668 + c.G*0.185562 + c.B*0.188199
	// Y = c.R*0.297344 + c.G*0.627376 + c.B*0.075279
	// Z = c.R*0.027031 + c.G*0.070690 + c.B*0.991179
}

func ExampleTransformFromXYZForXYYPrimaries_generateDisplayP3Matrix() {
	transform := ciexyz.TransformFromXYZForXYYPrimaries(
		displayp3.PrimaryRed,
		displayp3.PrimaryGreen,
		displayp3.PrimaryBlue,
		displayp3.StandardWhitePoint)

	fmt.Printf("R = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*2.493509 + c.Y*-0.931388 + c.Z*-0.402713
	// G = c.X*-0.829473 + c.Y*1.762631 + c.Z*0.023624
	// B = c.X*0.035851 + c.Y*-0.076184 + c.Z*0.957030
}

func ExampleTransformToXYZForXYYPrimaries_generateDisplayP3Matrix() {
	transform := ciexyz.TransformToXYZForXYYPrimaries(
		displayp3.PrimaryRed,
		displayp3.PrimaryGreen,
		displayp3.PrimaryBlue,
		displayp3.StandardWhitePoint)

	fmt.Printf("X = c.R*%f + c.G*%f + c.B*%f\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("Y = c.R*%f + c.G*%f + c.B*%f\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("Z = c.R*%f + c.G*%f + c.B*%f\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// X = c.R*0.486569 + c.G*0.265673 + c.B*0.198187
	// Y = c.R*0.228973 + c.G*0.691752 + c.B*0.079275
	// Z = c.R*0.000000 + c.G*0.045114 + c.B*1.043786
}

func ExampleTransformFromXYZForXYYPrimaries_generateSRGBMatrix() {
	transform := ciexyz.TransformFromXYZForXYYPrimaries(
		srgb.PrimaryRed,
		srgb.PrimaryGreen,
		srgb.PrimaryBlue,
		srgb.StandardWhitePoint)

	fmt.Printf("R = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%f + c.Y*%f + c.Z*%f\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*3.241004 + c.Y*-1.537399 + c.Z*-0.498616
	// G = c.X*-0.969224 + c.Y*1.875930 + c.Z*0.041554
	// B = c.X*0.055639 + c.Y*-0.204011 + c.Z*1.057149
}

func ExampleTransformToXYZForXYYPrimaries_generateSRGBMatrix() {
	transform := ciexyz.TransformToXYZForXYYPrimaries(
		srgb.PrimaryRed,
		srgb.PrimaryGreen,
		srgb.PrimaryBlue,
		srgb.StandardWhitePoint)

	fmt.Printf("X = c.R*%v + c.G*%v + c.B*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("Y = c.R*%v + c.G*%v + c.B*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("Z = c.R*%v + c.G*%v + c.B*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// X = c.R*0.41238652374145657 + c.G*0.35759149384555927 + c.B*0.18045052788799043
	// Y = c.R*0.21263680803732615 + c.G*0.7151829876911185 + c.B*0.07218020427155547
	// Z = c.R*0.019330619488581533 + c.G*0.11919711488225383 + c.B*0.9503727125209493
}

func ExampleTransformFromXYZForXYYPrimaries_generateProPhotoRGBMatrix() {
	transform := ciexyz.TransformFromXYZForXYYPrimaries(
		prophotorgb.PrimaryRed,
		prophotorgb.PrimaryGreen,
		prophotorgb.PrimaryBlue,
		prophotorgb.StandardWhitePoint)

	fmt.Printf("R = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*1.3459441751995702 + c.Y*-0.255601933365717 + c.Z*-0.05110784199842092
	// G = c.X*-0.5446048748563069 + c.Y*1.5081763487167865 + c.Z*0.020526475602742393
	// B = c.X*0 + c.Y*-0 + c.Z*1.2118446483846639
}

func ExampleTransformToXYZForXYYPrimaries_generateProPhotoRGBMatrix() {
	transform := ciexyz.TransformToXYZForXYYPrimaries(
		prophotorgb.PrimaryRed,
		prophotorgb.PrimaryGreen,
		prophotorgb.PrimaryBlue,
		prophotorgb.StandardWhitePoint)

	fmt.Printf("X = c.R*%v + c.G*%v + c.B*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("Y = c.R*%v + c.G*%v + c.B*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("Z = c.R*%v + c.G*%v + c.B*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// X = c.R*0.7976734029450476 + c.G*0.13518768157360342 + c.B*0.03135091585137471
	// Y = c.R*0.28804113269427045 + c.G*0.7118689212431865 + c.B*8.994606254312391e-05
	// Z = c.R*0 + c.G*0 + c.B*0.8251882791519165
}
