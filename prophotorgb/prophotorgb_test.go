package prophotorgb

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
