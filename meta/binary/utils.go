package binary

import (
	"io"
)

func ReadU16Big(r io.ByteReader) (uint16, error) {
	b1, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	b2, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	return uint16(b1)<<8 | uint16(b2), nil
}

func ReadU32Big(r io.ByteReader) (uint32, error) {
	b1, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	b2, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	b3, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	b4, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	return uint32(b1)<<24 | uint32(b2)<<16 | uint32(b3)<<8 | uint32(b4), nil
}

func ReadU64Big(r io.ByteReader) (uint64, error) {
	w1, err := ReadU32Big(r)
	if err != nil {
		return 0, err
	}

	w2, err := ReadU32Big(r)
	if err != nil {
		return 0, err
	}

	return uint64(w1)<<32 | uint64(w2), nil
}
