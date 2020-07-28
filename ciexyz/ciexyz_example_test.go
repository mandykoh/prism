package ciexyz_test

import (
	"fmt"
	"github.com/mandykoh/prism/adobergb"
	"github.com/mandykoh/prism/ciexyz"
)

func ExampleTransformFromXYZForPrimaries_generateAdobeRGBMatrix() {
	transform := ciexyz.TransformFromXYZForPrimaries(
		adobergb.PrimaryRedX, adobergb.PrimaryRedY,
		adobergb.PrimaryGreenX, adobergb.PrimaryGreenY,
		adobergb.PrimaryBlueX, adobergb.PrimaryBlueY,
		adobergb.StandardWhitePoint)

	fmt.Printf("R = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("G = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("B = c.X*%v + c.Y*%v + c.Z*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// R = c.X*2.0413690972656604 + c.Y*-0.5649464198330968 + c.Z*-0.3446944043036243
	// G = c.X*-0.969266003215008 + c.Y*1.8760107926266532 + c.Z*0.04155601636031898
	// B = c.X*0.013447387300642133 + c.Y*-0.11838974309780925 + c.Z*1.0154095783288708
}

func ExampleTransformToXYZForPrimaries_generateAdobeRGBMatrix() {
	transform := ciexyz.TransformToXYZForPrimaries(
		adobergb.PrimaryRedX, adobergb.PrimaryRedY,
		adobergb.PrimaryGreenX, adobergb.PrimaryGreenY,
		adobergb.PrimaryBlueX, adobergb.PrimaryBlueY,
		adobergb.StandardWhitePoint)

	fmt.Printf("X = c.R*%v + c.G*%v + c.B*%v\n", transform[0][0], transform[1][0], transform[2][0])
	fmt.Printf("Y = c.R*%v + c.G*%v + c.B*%v\n", transform[0][1], transform[1][1], transform[2][1])
	fmt.Printf("Z = c.R*%v + c.G*%v + c.B*%v\n", transform[0][2], transform[1][2], transform[2][2])

	// Output:
	// X = c.R*0.5767308538590208 + c.G*0.1855539559355801 + c.B*0.18818516090852389
	// Y = c.R*0.29737684652105756 + c.G*0.6273490891155328 + c.B*0.07527406436340955
	// Z = c.R*0.027034258774641568 + c.G*0.07068722130879249 + c.B*0.9911085141182259
}
