package autometa

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
)

func TestLoad(t *testing.T) {

	t.Run("returns original image data when format is unrecognised", func(t *testing.T) {
		randomBytes := make([]byte, 16)
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}

		input := bytes.NewReader(randomBytes)

		md, stream, err := Load(input)

		if err == nil {
			t.Fatalf("Expected error but succeeded")
		}
		if expected, actual := "unrecognised image format", err.Error(); expected != actual {
			t.Errorf("Expected error '%s' but was '%s'", expected, actual)
		}

		if md != nil {
			t.Errorf("Expected no metadata to be returned but was %+v", md)
		}

		returnedBytes, err := io.ReadAll(stream)
		if err != nil {
			t.Errorf("Expected to be able to read %d bytes from returned stream but got error: %v", len(randomBytes), err)
		}
		if expected, actual := len(randomBytes), len(returnedBytes); expected != actual {
			t.Fatalf("Expected returned stream to contain %d bytes but found %d", expected, actual)
		}

		if bytes.Compare(randomBytes, returnedBytes) != 0 {
			t.Errorf("Expected returned stream to contain original image data but was different.\n\nExpected:%v\nActual:%v\n", randomBytes, returnedBytes)
		}
	})
}
