package ciexyz

import (
	"github.com/mandykoh/prism/cielab"
	"math"
	"testing"
)

func TestColor(t *testing.T) {

	t.Run("ColorFromLAB()", func(t *testing.T) {

		t.Run("returns correct results", func(t *testing.T) {
			cases := []struct {
				WhitePoint  Color
				InputLAB    cielab.Color
				ExpectedXYZ Color
			}{
				// RGB(0, 0, 0)
				{WhitePoint: WhitePointD50, InputLAB: cielab.Color{L: 0, A: 0, B: 0, Alpha: 1}, ExpectedXYZ: Color{0, 0, 0, 1}},
				{WhitePoint: WhitePointD65, InputLAB: cielab.Color{L: 0, A: 0, B: 0, Alpha: 1}, ExpectedXYZ: Color{0, 0, 0, 1}},

				// RGB(255, 0, 0)
				{WhitePoint: WhitePointD50, InputLAB: cielab.Color{L: 54.29, A: 80.81, B: 69.89, Alpha: 1}, ExpectedXYZ: Color{0.4361, 0.2225, 0.0139, 1}},
				{WhitePoint: WhitePointD65, InputLAB: cielab.Color{L: 53.24, A: 80.09, B: 67.20, Alpha: 1}, ExpectedXYZ: Color{0.4125, 0.2127, 0.0193, 1}},

				// RGB(127, 127, 127)
				{WhitePoint: WhitePointD50, InputLAB: cielab.Color{L: 53.19, A: 0, B: 0, Alpha: 1}, ExpectedXYZ: Color{0.2046, 0.2122, 0.1751, 1}},
				{WhitePoint: WhitePointD65, InputLAB: cielab.Color{L: 53.19, A: 0, B: 0, Alpha: 1}, ExpectedXYZ: Color{0.2017, 0.2122, 0.2311, 1}},

				// RGB(0, 255, 255)
				{WhitePoint: WhitePointD50, InputLAB: cielab.Color{L: 90.67, A: -50.67, B: -14.96, Alpha: 1}, ExpectedXYZ: Color{0.5281, 0.7775, 0.8113, 1}},
				{WhitePoint: WhitePointD65, InputLAB: cielab.Color{L: 91.11, A: -48.09, B: -14.13, Alpha: 1}, ExpectedXYZ: Color{0.5380, 0.7873, 1.0695, 1}},

				// RGB(255, 255, 255)
				{WhitePoint: WhitePointD50, InputLAB: cielab.Color{L: 100, A: 0, B: 0, Alpha: 1}, ExpectedXYZ: Color{0.9642, 1.0, 0.8252, 1}},
				{WhitePoint: WhitePointD65, InputLAB: cielab.Color{L: 100, A: 0, B: 0, Alpha: 1}, ExpectedXYZ: Color{0.9505, 1.0, 1.0888, 1}},
			}

			for _, c := range cases {
				expected, actual := c.ExpectedXYZ, ColorFromLAB(c.InputLAB, c.WhitePoint)
				if math.Abs(float64(expected.X)-float64(actual.X)) > 0.01 ||
					math.Abs(float64(expected.Y)-float64(actual.Y)) > 0.01 ||
					math.Abs(float64(expected.Z)-float64(actual.Z)) > 0.01 ||
					math.Abs(float64(expected.A)-float64(actual.A)) > 0.01 {

					t.Errorf("Expected %+v with white point %+v to convert to %+v but was %+v", c.InputLAB, c.WhitePoint, expected, actual)
				}
			}
		})
	})

	t.Run("ToLAB()", func(t *testing.T) {

		t.Run("returns correct results", func(t *testing.T) {
			cases := []struct {
				WhitePoint  Color
				InputXYZ    Color
				ExpectedLAB cielab.Color
			}{
				// RGB(0, 0, 0)
				{WhitePoint: WhitePointD50, InputXYZ: Color{0, 0, 0, 1}, ExpectedLAB: cielab.Color{L: 0, A: 0, B: 0, Alpha: 1}},
				{WhitePoint: WhitePointD65, InputXYZ: Color{0, 0, 0, 1}, ExpectedLAB: cielab.Color{L: 0, A: 0, B: 0, Alpha: 1}},

				// RGB(255, 0, 0)
				{WhitePoint: WhitePointD50, InputXYZ: Color{0.4361, 0.2225, 0.0139, 1}, ExpectedLAB: cielab.Color{L: 54.291, A: 80.825, B: 69.922, Alpha: 1}},
				{WhitePoint: WhitePointD65, InputXYZ: Color{0.4125, 0.2127, 0.0193, 1}, ExpectedLAB: cielab.Color{L: 53.244, A: 80.093, B: 67.239, Alpha: 1}},

				// RGB(127, 127, 127)
				{WhitePoint: WhitePointD50, InputXYZ: Color{0.2046, 0.2122, 0.1751, 1}, ExpectedLAB: cielab.Color{L: 53.19, A: 0, B: 0, Alpha: 1}},
				{WhitePoint: WhitePointD65, InputXYZ: Color{0.2017, 0.2122, 0.2311, 1}, ExpectedLAB: cielab.Color{L: 53.19, A: 0, B: 0, Alpha: 1}},

				// RGB(0, 255, 255)
				{WhitePoint: WhitePointD50, InputXYZ: Color{0.5281, 0.7775, 0.8113, 1}, ExpectedLAB: cielab.Color{L: 90.666, A: -50.675, B: -14.972, Alpha: 1}},
				{WhitePoint: WhitePointD65, InputXYZ: Color{0.5380, 0.7873, 1.0695, 1}, ExpectedLAB: cielab.Color{L: 91.11, A: -48.09, B: -14.13, Alpha: 1}},

				// RGB(255, 255, 255)
				{WhitePoint: WhitePointD50, InputXYZ: Color{0.9642, 1.0, 0.8252, 1}, ExpectedLAB: cielab.Color{L: 100, A: 0, B: 0, Alpha: 1}},
				{WhitePoint: WhitePointD65, InputXYZ: Color{0.9505, 1.0, 1.0888, 1}, ExpectedLAB: cielab.Color{L: 100, A: 0, B: 0, Alpha: 1}},
			}

			for _, c := range cases {
				expected, actual := c.ExpectedLAB, c.InputXYZ.ToLAB(c.WhitePoint)
				if math.Abs(float64(expected.L)-float64(actual.L)) > 0.01 ||
					math.Abs(float64(expected.A)-float64(actual.A)) > 0.01 ||
					math.Abs(float64(expected.B)-float64(actual.B)) > 0.01 ||
					math.Abs(float64(expected.Alpha)-float64(actual.Alpha)) > 0.01 {

					t.Errorf("Expected %+v with white point %+v to convert to %+v but was %+v", c.InputXYZ, c.WhitePoint, expected, actual)
				}
			}
		})
	})
}
