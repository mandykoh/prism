package webpmeta

import (
	"bytes"
	"testing"

	"github.com/mandykoh/prism/meta/binary"
)

func TestExtractMetadata(t *testing.T) {
	t.Run("returns error with missing RIFF signature", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte("NOT A RIFF SIGNATURE"))

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "missing RIFF header", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns error with missing WEBP signature", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte("RIFF....NOTP"))

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "not a WEBP file", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns error if basic metadata is not found", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte("RIFF\x28\x51\x04\x00WEBPVP8"))

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "unexpected EOF reading chunk type", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns metadata without ICC profile if an ICC chunk is not present", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte("RIFF\xc0Z\x04\x00WEBPVP8X\x0a\x00\x00\x00\x14\x00\x00\x00\xaf\x04\x00\xaf\x04\x00"))

		md, err := extractMetadata(data)

		if err != nil {
			t.Fatalf("Expected success but got error: %v", err)
		}

		if md == nil {
			t.Errorf("Expected metdata but got none")
		} else {
			iccData, iccErr := md.ICCProfileData()

			if iccErr != nil {
				t.Errorf("Expected no ICC profile error but got: %v", iccErr)
			}
			if iccData != nil {
				t.Errorf("Expected no ICC profile but got one")
			}

			if expected, actual := uint32(1200), md.PixelWidth; expected != actual {
				t.Errorf("Expected image width of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(1200), md.PixelHeight; expected != actual {
				t.Errorf("Expected image height of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(8), md.BitsPerComponent; expected != actual {
				t.Errorf("Expected image bits per component of %d but got %d", expected, actual)
			}
		}
	})

	t.Run("returns all metadata", func(t *testing.T) {
		data := &bytes.Buffer{}
		iccProfileData := []byte{1, 2, 3, 4}
		data.Write([]byte("RIFF\xc0Z\x04\x00WEBPVP8X\x0a\x00\x00\x00\x34\x00\x00\x00\xaf\x04\x00\xaf\x04\x00"))
		data.Write(chunkTypeICCP[:])
		binary.WriteU32Little(data, uint32(len(iccProfileData)))
		data.Write(iccProfileData)

		md, err := extractMetadata(data)

		if err != nil {
			t.Fatalf("Expected success but got error: %v", err)
		}

		if md == nil {
			t.Errorf("Expected metdata but got none")
		} else {
			iccData, iccErr := md.ICCProfileData()

			if iccErr != nil {
				t.Errorf("Expected ICC profile data but got error: %v", iccErr)
			}

			if iccData == nil {
				t.Errorf("Expected an ICC profile but got none")
			} else {
				if expected, actual := len(iccProfileData), len(iccData); expected != actual {
					t.Errorf("Expected %d bytes of ICC profile data but got %d", expected, actual)
				}
				for i := range iccProfileData {
					if expected, actual := iccProfileData[i], iccData[i]; expected != actual {
						t.Fatalf("Expected ICC profile data %v but got %v", iccProfileData, iccData)
					}
				}
			}

			if expected, actual := uint32(1200), md.PixelWidth; expected != actual {
				t.Errorf("Expected image width of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(1200), md.PixelHeight; expected != actual {
				t.Errorf("Expected image height of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(8), md.BitsPerComponent; expected != actual {
				t.Errorf("Expected image bits per component of %d but got %d", expected, actual)
			}
		}
	})
}
