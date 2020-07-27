package srgb

import "testing"

func TestConvertLinearTo8Bit(t *testing.T) {

	t.Run("clips linear values to between 0 and 1", func(t *testing.T) {
		if expected, actual := ConvertLinearTo8Bit(0), ConvertLinearTo8Bit(-0.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
		if expected, actual := ConvertLinearTo8Bit(1), ConvertLinearTo8Bit(1.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
	})
}

func TestFrom8BitToLinear(t *testing.T) {

	t.Run("provides values for all possible 8-bit inputs", func(t *testing.T) {
		for i := 0; i < 256; i++ {
			if expected, actual := Convert8BitToLinear(uint8(i)), From8BitToLinear(uint8(i)); expected != actual {
				t.Errorf("Expected converted value to be %v but was %v", expected, actual)
			}
		}
	})
}

func TestFromLinearTo8Bit(t *testing.T) {

	t.Run("clips linear values to between 0 and 1", func(t *testing.T) {
		if expected, actual := ConvertLinearTo8Bit(0), FromLinearTo8Bit(-0.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
		if expected, actual := ConvertLinearTo8Bit(1), FromLinearTo8Bit(1.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
	})
}
