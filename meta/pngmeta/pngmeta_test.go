package pngmeta

import (
	"bytes"
	"compress/zlib"
	"github.com/mandykoh/prism/meta/binary"
	"testing"
)

func TestExtractMetadata(t *testing.T) {

	compressICCProfileData := func(data []byte) []byte {
		compressedICCProfileData := &bytes.Buffer{}
		zWriter := zlib.NewWriter(compressedICCProfileData)
		_, err := zWriter.Write(data)
		if err != nil {
			panic(err)
		}
		err = zWriter.Close()
		if err != nil {
			panic(err)
		}

		return compressedICCProfileData.Bytes()
	}

	writeICCProfileChunk := func(dst *bytes.Buffer, compressedICCData []byte) {
		profileName := []byte("SomeProfile")

		_ = binary.WriteU32Big(dst, uint32(len(profileName)+2+len(compressedICCData)))
		dst.Write(chunkTypeiCCP[:])
		dst.Write(profileName)
		dst.WriteByte(0x00)
		dst.WriteByte(0x00)
		dst.Write(compressedICCData)

		dummyCRC := uint32(0)
		_ = binary.WriteU32Big(dst, dummyCRC)
	}

	t.Run("returns error with invalid PNG signature", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write([]byte("NOT A PNG SIGNATURE"))

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "invalid PNG signature", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns error if basic metadata is not found", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write(pngSignature[:])

		_, err := extractMetadata(data)

		if err == nil {
			t.Errorf("Expected error but succeeded")
		} else if expected, actual := "no metadata found", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but got '%s'", expected, actual)
		}
	})

	t.Run("returns metadata without ICC profile if an ICC chunk is not present", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write(pngSignature[:])

		_ = binary.WriteU32Big(data, 13)
		data.Write(chunkTypeIHDR[:])
		headerData := [13]byte{0, 0, 0, 15, 0, 0, 0, 16, 8}
		data.Write(headerData[:])
		dummyCRC := uint32(0)
		_ = binary.WriteU32Big(data, dummyCRC)

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

			if expected, actual := uint32(15), md.PixelWidth; expected != actual {
				t.Errorf("Expected image width of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(16), md.PixelHeight; expected != actual {
				t.Errorf("Expected image height of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(8), md.BitsPerComponent; expected != actual {
				t.Errorf("Expected image bits per component of %d but got %d", expected, actual)
			}
		}
	})

	t.Run("returns metadata without ICC profile if the ICC data can't be decompressed", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write(pngSignature[:])

		_ = binary.WriteU32Big(data, 13)
		data.Write(chunkTypeIHDR[:])
		headerData := [13]byte{0, 0, 0, 15, 0, 0, 0, 16, 8}
		data.Write(headerData[:])
		dummyCRC := uint32(0)
		_ = binary.WriteU32Big(data, dummyCRC)

		invalidICCProfileData := []byte("NOT COMPRESSED ICC PROFLE DATA")
		writeICCProfileChunk(data, invalidICCProfileData)

		md, err := extractMetadata(data)

		if err != nil {
			t.Fatalf("Expected success but got error: %v", err)
		}

		if md == nil {
			t.Errorf("Expected metdata but got none")
		} else {
			iccData, iccErr := md.ICCProfileData()

			if iccData != nil {
				t.Errorf("Expected no ICC profile but got one")
			}

			if iccErr == nil {
				t.Errorf("Expected ICC profile error but got none")
			} else if expected, actual := "zlib: invalid header", iccErr.Error(); expected != actual {
				t.Errorf("Expected ICC profile error '%s' but got '%s'", expected, actual)
			}

			if expected, actual := uint32(15), md.PixelWidth; expected != actual {
				t.Errorf("Expected image width of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(16), md.PixelHeight; expected != actual {
				t.Errorf("Expected image height of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(8), md.BitsPerComponent; expected != actual {
				t.Errorf("Expected image bits per component of %d but got %d", expected, actual)
			}
		}
	})

	t.Run("returns all metadata", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write(pngSignature[:])

		_ = binary.WriteU32Big(data, 13)
		data.Write(chunkTypeIHDR[:])
		headerData := [13]byte{0, 0, 0, 15, 0, 0, 0, 16, 8}
		data.Write(headerData[:])
		dummyCRC := uint32(0)
		_ = binary.WriteU32Big(data, dummyCRC)

		iccProfileData := []byte{1, 2, 3, 4}
		writeICCProfileChunk(data, compressICCProfileData(iccProfileData))

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

			if expected, actual := uint32(15), md.PixelWidth; expected != actual {
				t.Errorf("Expected image width of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(16), md.PixelHeight; expected != actual {
				t.Errorf("Expected image height of %d but got %d", expected, actual)
			}
			if expected, actual := uint32(8), md.BitsPerComponent; expected != actual {
				t.Errorf("Expected image bits per component of %d but got %d", expected, actual)
			}
		}
	})

	t.Run("stops reading after all interesting metadata has been found", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write(pngSignature[:])

		_ = binary.WriteU32Big(data, 13)
		data.Write(chunkTypeIHDR[:])
		headerData := [13]byte{0, 0, 0, 16, 0, 0, 0, 16, 8}
		data.Write(headerData[:])
		dummyCRC := uint32(0)
		_ = binary.WriteU32Big(data, dummyCRC)

		iccProfileData := []byte{1, 2, 3, 4}
		writeICCProfileChunk(data, compressICCProfileData(iccProfileData))

		_ = binary.WriteU32Big(data, 4)
		data.Write(chunkTypeIDAT[:])
		imageData := []byte{5, 6, 7, 8}
		data.Write(imageData)
		_ = binary.WriteU32Big(data, dummyCRC)

		_, err := extractMetadata(data)

		if err != nil {
			t.Fatalf("Expected success but got error: %v", err)
		}

		ch, err := readChunkHeader(data)
		if err != nil {
			t.Fatalf("Expected chunk data to follow metadata but got error: %v", err)
		} else {
			if expected, actual := (chunkHeader{4, chunkTypeIDAT}), ch; expected != actual {
				t.Fatalf("Expected chunk %v but got %v", expected, actual)
			}
		}
	})

	t.Run("stops reading at IDAT chunk header", func(t *testing.T) {
		data := &bytes.Buffer{}
		data.Write(pngSignature[:])

		_ = binary.WriteU32Big(data, 13)
		data.Write(chunkTypeIHDR[:])
		headerData := [13]byte{}
		data.Write(headerData[:])
		dummyCRC := uint32(0)
		_ = binary.WriteU32Big(data, dummyCRC)

		_ = binary.WriteU32Big(data, 999)
		data.Write(chunkTypeIDAT[:])
		imageData := []byte{1, 2, 3, 4}
		data.Write(imageData)

		_, err := extractMetadata(data)

		if err != nil {
			t.Fatalf("Expected success but got error: %v", err)
		}

		b := [4]byte{}
		n, err := data.Read(b[:])
		if err != nil {
			t.Fatalf("Expected data to be available after IDAT header but got error: %v", err)
		}
		if n < 4 {
			t.Fatalf("Expected IDAT data but got %v", b[:n])
		}
		if expected, actual := [4]byte{1, 2, 3, 4}, b; expected != actual {
			t.Errorf("Expected bytes %v to be available after IDAT header but got %v", expected, actual)
		}
	})
}
