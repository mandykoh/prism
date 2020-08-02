package jpegmeta

import (
	"github.com/mandykoh/prism/meta/binary"
)

type segmentReader struct {
	reader             binary.Reader
	inEntropyCodedData bool
}

func (sr *segmentReader) ReadSegment() (segment, error) {

	if sr.inEntropyCodedData {
		for {
			b, err := sr.reader.ReadByte()
			if err != nil {
				return segment{}, err
			}

			if b == 0xFF {
				b, err = sr.reader.ReadByte()
				if err != nil {
					return segment{}, err
				}

				if b != 0x00 {
					seg, err := makeSegment(b, sr.reader)
					if err != nil {
						return segment{}, err
					}

					sr.inEntropyCodedData = seg.Marker.Type == MarkerTypeStartOfScan ||
						(seg.Marker.Type >= MarkerTypeRestart0 && seg.Marker.Type <= MarkerTypeRestart7)

					return seg, err
				}
			}
		}
	}

	seg, err := readSegment(sr.reader)
	if err != nil {
		return seg, err
	}

	sr.inEntropyCodedData = seg.Marker.Type == MarkerTypeStartOfScan

	return seg, nil
}

func NewSegmentReader(r binary.Reader) *segmentReader {
	return &segmentReader{
		reader: r,
	}
}
