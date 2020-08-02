package icc

import (
	"fmt"
)

type TagTable struct {
	entries map[Signature][]byte
}

func (t *TagTable) add(sig Signature, data []byte) {
	t.entries[sig] = data
}

func (t *TagTable) getProfileDescription(ver Version) (string, error) {
	data := t.entries[DescSignature]

	switch ver.Major {

	case 2:
		desc, err := parseTextDescription(data)
		if err != nil {
			return "", err
		}
		return desc.ASCII, nil

	case 4:
		mluc, err := parseMultiLocalisedUnicode(data)
		if err != nil {
			return "", err
		}
		if enUS := mluc.getStringForLanguage([2]byte{'e', 'n'}); enUS != "" {
			return enUS, nil
		}
		return mluc.getAnyString(), nil

	default:
		return "", fmt.Errorf("unknown profile major version (%d)", ver.Major)
	}
}

func emptyTagTable() TagTable {
	return TagTable{
		entries: make(map[Signature][]byte),
	}
}
