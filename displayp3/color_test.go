package displayp3

import (
	"github.com/mandykoh/prism/srgb"
	"image/color"
	"math"
	"testing"
)

func TestColor(t *testing.T) {

	t.Run("ColorFromRGBA()", func(t *testing.T) {

		t.Run("returns correct results for full alpha", func(t *testing.T) {
			for i := 0; i < 256; i++ {
				nrgba := color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
				expected, expectedAlpha := ColorFromNRGBA(nrgba)

				rgba := color.RGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
				actual, actualAlpha := ColorFromRGBA(rgba)

				if expected != actual {
					t.Errorf("Expected %+v to map to %+v but was %+v", rgba, expected, actual)
				}
				if math.Abs(float64(expectedAlpha)-float64(actualAlpha)) > 0.0001 {
					t.Errorf("Expected alpha %d to map to %v but was %v", rgba.A, expectedAlpha, actualAlpha)
				}
			}
		})

		t.Run("returns correct results for scaled alpha", func(t *testing.T) {
			for i := 0; i < 256; i++ {
				expectedAlpha := float32(i) / 255

				var expected Color
				if expectedAlpha > 0 {
					expected = Color{
						R: srgb.From8Bit(uint8(i)) / expectedAlpha,
						G: srgb.From8Bit(uint8(i)) / expectedAlpha,
						B: srgb.From8Bit(uint8(i)) / expectedAlpha,
					}
				}

				rgba := color.RGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: uint8(i)}
				actual, actualAlpha := ColorFromRGBA(rgba)

				if expected != actual {
					t.Errorf("Expected %+v to map to %+v but was %+v", rgba, expected, actual)
				}
				if math.Abs(float64(expectedAlpha)-float64(actualAlpha)) > 0.0001 {
					t.Errorf("Expected alpha %d to map to %v but was %v", rgba.A, expectedAlpha, actualAlpha)
				}
			}
		})
	})

	t.Run("ToRGBA()", func(t *testing.T) {

		t.Run("returns correct results for full alpha", func(t *testing.T) {
			for i := 0; i < 256; i++ {
				c, a := ColorFromNRGBA(color.NRGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255})

				nrgba := c.ToNRGBA(a)
				expected := color.RGBA{
					R: nrgba.R,
					G: nrgba.G,
					B: nrgba.B,
					A: nrgba.A,
				}

				actual := c.ToRGBA(a)

				if expected != actual {
					t.Errorf("Expected normalised %+v with alpha %v to map to %+v but was %+v", c, a, expected, actual)
				}
			}
		})

		t.Run("returns correct results for scaled alpha", func(t *testing.T) {
			for i := 0; i < 256; i++ {
				a := float32(i) / 255

				var expected color.RGBA
				if a > 0 {
					expected = color.RGBA{
						R: srgb.To8Bit(a * a),
						G: srgb.To8Bit(a * a),
						B: srgb.To8Bit(a * a),
						A: uint8(a * 255),
					}
				}

				c := Color{R: a, G: a, B: a}
				actual := c.ToRGBA(a)

				if expected != actual {
					t.Errorf("Expected normalised %+v with alpha %v to map to %+v but was %+v", c, a, expected, actual)
				}
			}
		})
	})
}
