package prophotorgb

import "testing"

func TestFrom8Bit(t *testing.T) {

	t.Run("provides linear values for all possible 8-bit inputs", func(t *testing.T) {
		for i := 0; i < 256; i++ {
			if expected, actual := float32(Convert8BitToLinear(uint8(i))), From8Bit(uint8(i)); expected != actual {
				t.Errorf("Expected converted value to be %v but was %v", expected, actual)
			}
		}
	})
}

func TestTo8Bit(t *testing.T) {

	t.Run("clips linear values to between 0 and 1", func(t *testing.T) {
		if expected, actual := ConvertLinearTo8Bit(0), To8Bit(-0.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
		if expected, actual := ConvertLinearTo8Bit(1), To8Bit(1.1); expected != actual {
			t.Errorf("Expected converted value to be %v but was %v", expected, actual)
		}
	})
}
