package jpegmeta

import (
	"bytes"
	"github.com/mandykoh/prism/meta/binary"
	"math/rand"
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

	writeICCProfileHeader := func(w *bytes.Buffer) {
		profileSize := uint32(rand.Int31())
		_ = binary.WriteU32Big(w, profileSize)

		_, _ = w.Write([]byte{'t', 'e', 's', 't'})                 // Preferred CMM
		_, _ = w.Write([]byte{4, 0, 0, 0})                         // Version
		_, _ = w.Write([]byte{'t', 'e', 's', 't'})                 // Device class
		_, _ = w.Write([]byte{'R', 'G', 'B', ' '})                 // Data colour space
		_, _ = w.Write([]byte{'X', 'Y', 'Z', ' '})                 // Profile connection space
		_, _ = w.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) // Creation date/time
		_, _ = w.Write([]byte{'a', 'c', 's', 'p'})                 // Profile signature
		_, _ = w.Write([]byte{'t', 'e', 's', 't'})                 // Primary platform
		_, _ = w.Write([]byte{0, 0, 0, 0})                         // Profile flags
		_, _ = w.Write([]byte{0, 0, 0, 0})                         // Device manufacturer
		_, _ = w.Write([]byte{0, 0, 0, 0})                         // Device model
		_, _ = w.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0})             // Device attributes
		_, _ = w.Write([]byte{0, 0, 0, 0})                         // Rendering intent
		_, _ = w.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) // PCS illuminant
		_, _ = w.Write([]byte{0, 0, 0, 0})                         // Profile creator

		profileID := [16]byte{}
		_, _ = w.Write(profileID[:])

		reservedBytes := [28]byte{}
		_, _ = w.Write(reservedBytes[:])
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

		iccProfile := &bytes.Buffer{}
		writeICCProfileHeader(iccProfile)
		writeICCProfileChunk(data, 1, 2, iccProfile.Bytes())

		data.Write([]byte{0xFF, byte(markerTypeStartOfScan), 0x00, 0x02})

		md, err := extractMetadata(data)

		if err != nil {
			t.Errorf("Expected success but got error: %v", err)
		} else {
			profile, err := md.ICCProfile()
			if err != nil {
				t.Errorf("Expected no ICC profile but got error: %v", err)
			} else if profile != nil {
				t.Errorf("Expected no ICC profile but got one")
			}
		}
	})

	t.Run("stops reading after all interesting metadata has been found", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte{0xFF, byte(markerTypeStartOfImage)})
		data.Write([]byte{0xFF, byte(markerTypeStartOfFrameBaseline), 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00})

		iccProfile := &bytes.Buffer{}
		writeICCProfileHeader(iccProfile)
		writeICCProfileChunk(data, 1, 1, iccProfile.Bytes())

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
