package jpegmeta

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mandykoh/prism/meta"
	"github.com/mandykoh/prism/meta/icc"
	"io"
)

var iccProfileIdentifier = []byte("ICC_PROFILE\x00")

// Load loads the metadata for a JPEG image stream.
//
// Only as much of the stream is consumed as necessary to extract the metadata;
// the returned stream contains a buffered copy of the consumed data such that
// reading from it will produce the same results as fully reading the input
// stream. This provides a convenient way to load the full image after loading
// the metadata.
//
// An error is returned if basic metadata could not be extracted. The returned
// stream still provides the full image data.
func Load(r io.Reader) (md *meta.Data, imgStream io.Reader, err error) {
	rewindBuffer := &bytes.Buffer{}
	tee := io.TeeReader(r, rewindBuffer)
	md, err = extractMetadata(tee)
	return md, io.MultiReader(rewindBuffer, r), err
}

func extractMetadata(r io.Reader) (*meta.Data, error) {
	md := &meta.Data{}
	segReader := NewSegmentReader(bufio.NewReader(r))

	var iccProfileChunks [][]byte

parseSegments:
	for {
		segment, err := segReader.ReadSegment()
		if err != nil {
			if err == io.EOF {
				return nil, fmt.Errorf("unexpected EOF")
			}
			return nil, err
		}

		switch segment.Marker.Type {

		case MarkerTypeStartOfFrameBaseline,
			MarkerTypeStartOfFrameProgressive:
			md.BitsPerComponent = int(segment.Data[0])
			md.PixelHeight = int(segment.Data[1])<<8 | int(segment.Data[2])
			md.PixelWidth = int(segment.Data[3])<<8 | int(segment.Data[4])

		case MarkerTypeStartOfScan:
			break parseSegments

		case MarkerTypeApp2:
			if len(segment.Data) < len(iccProfileIdentifier)+2 {
				continue
			}

			for i := range iccProfileIdentifier {
				if segment.Data[i] != iccProfileIdentifier[i] {
					continue
				}
			}

			chunkTotal := segment.Data[len(iccProfileIdentifier)+1]
			if iccProfileChunks == nil {
				iccProfileChunks = make([][]byte, chunkTotal)
			} else if int(chunkTotal) != len(iccProfileChunks) {
				return nil, fmt.Errorf("inconsistent ICC profile chunk count")
			}

			chunkNum := segment.Data[len(iccProfileIdentifier)]
			if chunkNum == 0 || int(chunkNum) > len(iccProfileChunks) {
				return nil, fmt.Errorf("invalid ICC profile chunk number")
			}
			iccProfileChunks[chunkNum-1] = segment.Data[len(iccProfileIdentifier)+2:]
		}
	}

	// No ICC profile
	if len(iccProfileChunks) == 0 {
		return md, nil
	}

	iccProfileData := bytes.Buffer{}
	for i := range iccProfileChunks {
		iccProfileData.Write(iccProfileChunks[i])
	}

	iccProfile, err := icc.NewProfileReader(bytes.NewReader(iccProfileData.Bytes())).ReadProfile()
	if err != nil {
		md.SetICCProfileErr(err)
	} else {
		md.SetICCProfile(iccProfile)
	}

	return md, nil
}
