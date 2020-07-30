package ciexyz_test

import (
	"fmt"
	"github.com/mandykoh/prism/adobergb"
	"github.com/mandykoh/prism/ciexyz"
	"github.com/mandykoh/prism/prophotorgb"
	"github.com/mandykoh/prism/srgb"
)

func ExampleTransformFromXYZForXYYPrimaries_generateAdobeRGBMatrix() {
	transform := ciexyz.TransformFromXYZForXYYPrimaries(
		adobergb.PrimaryRed,
		adobergb.PrimaryGreen,
		adobergb.PrimaryBlue,
		adobergb.StandardWhitePoint)

	fmt.Printf("R = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*2.0415913017647322 + c.Y*-0.5650078698012716 + c.Z*-0.34473195659062167
	// G = c.X*-0.9692242864995342 + c.Y*1.8759299885141114 + c.Z*0.04155424903337176
	// B = c.X*0.013446472278330708 + c.Y*-0.11838142234726094 + c.Z*1.01533754937275
}

func ExampleTransformToXYZForXYYPrimaries_generateAdobeRGBMatrix() {
	transform := ciexyz.TransformToXYZForXYYPrimaries(
		adobergb.PrimaryRed,
		adobergb.PrimaryGreen,
		adobergb.PrimaryBlue,
		adobergb.StandardWhitePoint)

	fmt.Printf("X = c.R*%v + c.G*%v + c.B*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("Y = c.R*%v + c.G*%v + c.B*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("Z = c.R*%v + c.G*%v + c.B*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// X = c.R*0.5766680793281725 + c.G*0.1855619421659935 + c.B*0.18819852398084014
	// Y = c.R*0.29734448781899253 + c.G*0.6273761097678748 + c.B*0.07527940241313279
	// Z = c.R*0.027031317880049893 + c.G*0.07069030664147563 + c.B*0.9911788223702592
}

func ExampleTransformFromXYZForXYYPrimaries_generateSRGBMatrix() {
	transform := ciexyz.TransformFromXYZForXYYPrimaries(
		srgb.PrimaryRed,
		srgb.PrimaryGreen,
		srgb.PrimaryBlue,
		srgb.StandardWhitePoint)

	fmt.Printf("R = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*3.241003600540255 + c.Y*-1.5373991710891957 + c.Z*-0.49861598312439226
	// G = c.X*-0.9692242864995344 + c.Y*1.8759299885141119 + c.Z*0.04155424903337176
	// B = c.X*0.05563936186796137 + c.Y*-0.20401108051523723 + c.Z*1.0571488385644063
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
