package meta

import (
	"bytes"
	"github.com/mandykoh/prism/meta/icc"
)

// Data represents the metadata for an image.
type Data struct {
	PixelWidth       uint32
	PixelHeight      uint32
	BitsPerComponent uint32
	ICCProfileData   []byte
}

// ICCProfile returns an extracted ICC profile from this metadata.
//
// An error is returned if the ICC profile could not be correctly parsed.
//
// If no profile data was found, nil is returned without an error.
func (md *Data) ICCProfile() (*icc.Profile, error) {
	if md.ICCProfileData == nil {
		return nil, nil
	}

	return icc.NewProfileReader(bytes.NewReader(md.ICCProfileData)).ReadProfile()
}
