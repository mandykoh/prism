package linear

import "testing"

func TestNormalisedTo8Bit(t *testing.T) {

	t.Run("clamps lower values to zero", func(t *testing.T) {
		if expected, actual := uint8(0), NormalisedTo8Bit(-0.1); expected != actual {
			t.Errorf("Expected negative value to be clamped to %d but was %d", expected, actual)
		}
	})

	t.Run("clamps higher values to 255", func(t *testing.T) {
		if expected, actual := uint8(255), NormalisedTo8Bit(1.1); expected != actual {
			t.Errorf("Expected high value to be clamped to %d but was %d", expected, actual)
		}
	})

	t.Run("rounds to nearest integer", func(t *testing.T) {
		if expected, actual := uint8(127), NormalisedTo8Bit(127.49/255.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
		if expected, actual := uint8(128), NormalisedTo8Bit(127.5/255.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
	})

	t.Run("rounding doesn't overflow", func(t *testing.T) {
		if expected, actual := uint8(255), NormalisedTo8Bit(254.9/255.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
	})
}

func TestNormalisedTo9Bit(t *testing.T) {

	t.Run("clamps lower values to zero", func(t *testing.T) {
		if expected, actual := uint16(0), NormalisedTo9Bit(-0.1); expected != actual {
			t.Errorf("Expected negative value to be clamped to %d but was %d", expected, actual)
		}
	})

	t.Run("clamps higher values to 511", func(t *testing.T) {
		if expected, actual := uint16(511), NormalisedTo9Bit(1.1); expected != actual {
			t.Errorf("Expected high value to be clamped to %d but was %d", expected, actual)
		}
	})

	t.Run("rounds to nearest integer", func(t *testing.T) {
		if expected, actual := uint16(255), NormalisedTo9Bit(255.49/511.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
		if expected, actual := uint16(256), NormalisedTo9Bit(255.5/511.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
	})

	t.Run("rounding doesn't overflow", func(t *testing.T) {
		if expected, actual := uint16(511), NormalisedTo9Bit(510.9/511.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
	})
}

func TestNormalisedTo16Bit(t *testing.T) {

	t.Run("clamps lower values to zero", func(t *testing.T) {
		if expected, actual := uint16(0), NormalisedTo16Bit(-0.1); expected != actual {
			t.Errorf("Expected negative value to be clamped to %d but was %d", expected, actual)
		}
	})

	t.Run("clamps higher values to 65535", func(t *testing.T) {
		if expected, actual := uint16(65535), NormalisedTo16Bit(1.1); expected != actual {
			t.Errorf("Expected high value to be clamped to %d but was %d", expected, actual)
		}
	})

	t.Run("rounds to nearest integer", func(t *testing.T) {
		if expected, actual := uint16(32767), NormalisedTo16Bit(32767.49/65535.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
		if expected, actual := uint16(32768), NormalisedTo16Bit(32767.5/65535.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
	})

	t.Run("rounding doesn't overflow", func(t *testing.T) {
		if expected, actual := uint16(65535), NormalisedTo16Bit(65534.9/65535.0); expected != actual {
			t.Errorf("Expected value to be rounded to %d but was %d", expected, actual)
		}
	})
}
