package jpegmeta

import (
	"fmt"
	"github.com/mandykoh/prism/meta/binary"
	"io"
)

var invalidMarker = marker{Type: MarkerTypeInvalid}

type marker struct {
	Type       markerType
	DataLength int
}

func makeMarker(mType byte, r io.ByteReader) (marker, error) {
	var length uint16
	switch mType {

	case
		byte(MarkerTypeRestart0),
		byte(MarkerTypeRestart1),
		byte(MarkerTypeRestart2),
		byte(MarkerTypeRestart3),
		byte(MarkerTypeRestart4),
		byte(MarkerTypeRestart5),
		byte(MarkerTypeRestart6),
		byte(MarkerTypeRestart7),
		byte(MarkerTypeStartOfImage),
		byte(MarkerTypeEndOfImage):

		length = 2

	case byte(MarkerTypeStartOfFrameBaseline),
		byte(MarkerTypeStartOfFrameProgressive),
		byte(MarkerTypeDefineHuffmanTable),
		byte(MarkerTypeStartOfScan),
		byte(MarkerTypeDefineQuantisationTable),
		byte(MarkerTypeDefineRestartInterval),
		byte(MarkerTypeApp0),
		byte(MarkerTypeApp1),
		byte(MarkerTypeApp2),
		byte(MarkerTypeApp3),
		byte(MarkerTypeApp4),
		byte(MarkerTypeApp5),
		byte(MarkerTypeApp6),
		byte(MarkerTypeApp7),
		byte(MarkerTypeApp8),
		byte(MarkerTypeApp9),
		byte(MarkerTypeApp10),
		byte(MarkerTypeApp11),
		byte(MarkerTypeApp12),
		byte(MarkerTypeApp13),
		byte(MarkerTypeApp14),
		byte(MarkerTypeApp15),
		byte(MarkerTypeComment):

		var err error
		length, err = binary.ReadU16Big(r)
		if err != nil {
			return invalidMarker, err
		}

	default:
		return invalidMarker, fmt.Errorf("unrecognised marker type %0x", mType)
	}

	return marker{
		Type:       markerType(mType),
		DataLength: int(length) - 2,
	}, nil
}

func readMarker(r io.ByteReader) (marker, error) {
	b, err := r.ReadByte()
	if err != nil {
		return invalidMarker, err
	}

	if b != 0xff {
		return invalidMarker, fmt.Errorf("invalid marker identifier %0x", b)
	}

	b, err = r.ReadByte()
	if err != nil {
		return invalidMarker, err
	}

	return makeMarker(b, r)
}
