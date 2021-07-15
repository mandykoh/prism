package icc

import (
	"bytes"
	"fmt"
	"unicode/utf16"

	"github.com/mandykoh/prism/meta/binary"
)

type MultiLocalisedUnicode struct {
	entriesByLanguageCountry map[[2]byte]map[[2]byte]string
}

func (mluc *MultiLocalisedUnicode) getAnyString() string {
	for _, country := range mluc.entriesByLanguageCountry {
		for _, s := range country {
			return s
		}
	}
	return ""
}

func (mluc *MultiLocalisedUnicode) getString(language [2]byte, country [2]byte) string {
	countries, ok := mluc.entriesByLanguageCountry[language]
	if !ok {
		return ""
	}

	return countries[country]
}

func (mluc *MultiLocalisedUnicode) getStringForLanguage(language [2]byte) string {
	for _, s := range mluc.entriesByLanguageCountry[language] {
		return s
	}
	return ""
}

func (mluc *MultiLocalisedUnicode) setString(language [2]byte, country [2]byte, text string) {
	countries, ok := mluc.entriesByLanguageCountry[language]
	if !ok {
		countries = map[[2]byte]string{
			country: text,
		}
		mluc.entriesByLanguageCountry[language] = countries

	} else {
		countries[country] = text
	}
}

func parseMultiLocalisedUnicode(data []byte) (MultiLocalisedUnicode, error) {
	result := MultiLocalisedUnicode{
		entriesByLanguageCountry: make(map[[2]byte]map[[2]byte]string),
	}

	reader := bytes.NewReader(data)

	// Parse the signature field:
	sig, err := binary.ReadU32Big(reader)
	if err != nil {
		return result, err
	}
	s := Signature(sig)

	// Some v4 profiles provide backwards compatibility for v2 readers by adding a v2 ASCII signature.
	// Per the spec, we should map this from ASCII to an en-US localized UTF-16 string.
	var hasDescSignature = false
	if s == DescSignature {
		hasDescSignature = true

		// Parse the v2 text description:
		desc, err := parseTextDescriptionFromReader(reader)
		if err != nil {
			return result, err
		}

		language := [2]byte{'e', 'n'}
		country := [2]byte{'u', 's'}
		result.setString(language, country, desc.ASCII)

		// Parse the next signature field:
		sig, err = binary.ReadU32Big(reader)
		if err != nil {
			return result, err
		}
		s = Signature(sig)
	}

	// If this is not a multi-localized unicode signature, return an error.
	if s != MultiLocalisedUnicodeSignature {
		if hasDescSignature {
			// If we had a desc signature just return that since it's our only tag.
			return result, nil
		}

		return result, fmt.Errorf("expected %v but got %v", MultiLocalisedUnicodeSignature, s)
	}

	// Reserved field
	_, err = binary.ReadU32Big(reader)
	if err != nil {
		return result, err
	}

	recordCount, err := binary.ReadU32Big(reader)
	if err != nil {
		return result, err
	}

	recordSize, err := binary.ReadU32Big(reader)
	if err != nil {
		return result, err
	}

	for i := uint32(0); i < recordCount; i++ {

		language := [2]byte{}
		n, err := reader.Read(language[:])
		if err != nil {
			return result, err
		}
		if n < len(language) {
			return result, fmt.Errorf("unexpected eof when reading language code")
		}

		country := [2]byte{}
		n, err = reader.Read(country[:])
		if err != nil {
			return result, err
		}
		if n < len(country) {
			return result, fmt.Errorf("unexpected eof when reading country code")
		}

		stringLength, err := binary.ReadU32Big(reader)
		if err != nil {
			return result, err
		}

		stringOffset, err := binary.ReadU32Big(reader)
		if err != nil {
			return result, err
		}

		if uint64(stringOffset+stringLength) > uint64(len(data)) {
			return result, fmt.Errorf("record exceeds tag data length")
		}

		recordStringBytes := data[stringOffset : stringOffset+stringLength]
		recordStringUTF16 := make([]uint16, len(recordStringBytes)/2)
		for j := 0; j < len(recordStringUTF16); j++ {
			recordStringUTF16[j], err = binary.ReadU16Big(reader)
			if err != nil {
				return result, err
			}
		}
		result.setString(language, country, string(utf16.Decode(recordStringUTF16)))

		// Skip to next record
		for j := uint32(12); j < recordSize; j++ {
			_, err := reader.ReadByte()
			if err != nil {
				return result, err
			}
		}
	}

	return result, nil
}
