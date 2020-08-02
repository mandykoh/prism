package jpegmeta

import (
	"bytes"
	"testing"
)

func TestExtractMetadata(t *testing.T) {

	t.Run("returns error with no start-of-frame segment", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeEndOfImage)})

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected an error but succeeded")
		} else if expected, actual := "no metadata found", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("stops reading at start-of-scan segment", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})
		data.Write([]byte{0xFF, byte(markerTypeStartOfScan), 0x00, 0x02})
		data.Write([]byte{0xFF, byte(markerTypeEndOfImage)})

		_, err := extractMetadata(data)

		if err != nil {
			t.Fatalf("Expected success but got error: %v", err)
		}

		b := [2]byte{}
		n, err := data.Read(b[:])
		if err != nil {
			t.Fatalf("Expected data to be available after start-of-scan but got error: %v", err)
		}
		if n < 2 {
			t.Fatalf("Expected end-of-image marker but got %v", b[:n])
		}
		if expected, actual := [2]byte{0xFF, byte(markerTypeEndOfImage)}, b; expected != actual {
			t.Errorf("Expected bytes %v to be available after start-of-scan but got %v", expected, actual)
		}
	})
}
