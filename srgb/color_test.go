package srgb

import (
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
				nrgba := color.NRGBA{R: 255, G: 255, B: 255, A: uint8(i)}
				if i == 0 {
					nrgba = color.NRGBA{R: 0, G: 0, B: 0, A: 0}
				}
				expected, expectedAlpha := ColorFromNRGBA(nrgba)

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
				c, a := ColorFromNRGBA(color.NRGBA{R: 255, G: 255, B: 255, A: uint8(i)})

				nrgba := c.ToNRGBA(a)
				var expected color.RGBA
				if nrgba.A == 0 {
					expected = color.RGBA{R: 0, G: 0, B: 0, A: 0}
				} else {
					expected = color.RGBA{
						R: uint8((int(nrgba.R) * int(nrgba.A)) / 255),
						G: uint8((int(nrgba.G) * int(nrgba.A)) / 255),
						B: uint8((int(nrgba.B) * int(nrgba.A)) / 255),
						A: nrgba.A,
					}
				}

				actual := c.ToRGBA(a)

				if expected != actual {
					t.Errorf("Expected normalised %+v with alpha %v to map to %+v but was %+v", c, a, expected, actual)
				}
			}
		})
	})
}
