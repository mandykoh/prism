package jpegmeta

import (
	"bytes"
	"testing"
)

func TestExtractMetadata(t *testing.T) {

	writeICCProfileChunk := func(dest *bytes.Buffer, chunkNum, chunkTotal byte, chunkData []byte) {
		dest.Write([]byte{0xFF, byte(markerTypeApp2)})
		chunkLength := len(iccProfileIdentifier) + 4 + len(chunkData)
		dest.Write([]byte{byte(chunkLength >> 8 & 0xFF), byte(chunkLength & 0xFF)})
		dest.Write(iccProfileIdentifier)
		dest.Write([]byte{chunkNum, chunkTotal})
		dest.Write(chunkData)
	}

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

	t.Run("returns error if ICC chunk number is higher than total chunks", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})
		writeICCProfileChunk(data, 2, 1, nil)

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "invalid ICC profile chunk number", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns error if subsequent ICC chunks specify different total chunks", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})
		writeICCProfileChunk(data, 1, 2, nil)
		writeICCProfileChunk(data, 2, 3, nil)

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "inconsistent ICC profile chunk count", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns error if an ICC chunk is duplicated", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})
		writeICCProfileChunk(data, 1, 3, nil)
		writeICCProfileChunk(data, 1, 3, nil)

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "duplicated ICC profile chunk", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns metadata without ICC profile if an ICC chunk is missing", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})

		iccProfileData := []byte{0, 1, 2, 3, 4, 5}
		writeICCProfileChunk(data, 1, 2, iccProfileData)

		data.Write([]byte{0xFF, byte(markerTypeStartOfScan), 0x00, 0x02})

		md, err := extractMetadata(data)

		if err != nil {
			t.Errorf("Expected success but got error: %v", err)
		} else if md.ICCProfileData != nil {
			t.Errorf("Expected no ICC profile but got one")
		}
	})

	t.Run("extracts ICC profile data from all ICC profile chunks", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})

		iccProfileData1 := []byte{0, 1, 2, 3, 4, 5}
		writeICCProfileChunk(data, 2, 2, iccProfileData1)

		iccProfileData2 := []byte{6, 7, 8, 9, 10, 11}
		writeICCProfileChunk(data, 1, 2, iccProfileData2)

		data.Write([]byte{0xFF, byte(markerTypeStartOfScan), 0x00, 0x02})

		md, err := extractMetadata(data)

		if err != nil {
			t.Errorf("Expected success but got error: %v", err)
		}

		if md.ICCProfileData == nil {
			t.Errorf("Expected ICC profile data to be extracted but got none")
		} else {
			if expected, actual := len(iccProfileData1)+len(iccProfileData2), len(md.ICCProfileData); expected != actual {
				t.Fatalf("Expected %d bytes of ICC profile data but got %d", expected, actual)
			}

			for i := range iccProfileData2 {
				if expected, actual := iccProfileData2[i], md.ICCProfileData[i]; expected != actual {
					t.Fatalf("Expected ICC profile byte %02x but got %02x", expected, actual)
				}
			}
			for i := range iccProfileData1 {
				if expected, actual := iccProfileData1[i], md.ICCProfileData[len(iccProfileData2)+i]; expected != actual {
					t.Fatalf("Expected ICC profile byte %02x but got %02x", expected, actual)
				}
			}
		}
	})

	t.Run("stops reading after all interesting metadata has been found", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})

		iccProfileData := []byte{0, 1, 2, 3, 4, 5}
		writeICCProfileChunk(data, 1, 1, iccProfileData)

		data.Write([]byte{0xFF, byte(markerTypeStartOfScan), 0x00, 0x02})

		_, err := extractMetadata(data)

		if err != nil {
			t.Errorf("Expected success but got error: %v", err)
		}

		b := [2]byte{}
		n, err := data.Read(b[:])
		if err != nil {
			t.Fatalf("Expected data to be available after ICC profile but got error: %v", err)
		}
		if n < 2 {
			t.Fatalf("Expected start-of-scan marker but got %v", b[:n])
		}
		if expected, actual := [2]byte{0xFF, byte(markerTypeStartOfScan)}, b; expected != actual {
			t.Errorf("Expected bytes %v to be available after ICC profile but got %v", expected, actual)
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
