package meta

import (
	"github.com/mandykoh/prism/meta/icc"
)

type Data struct {
	PixelWidth       int
	PixelHeight      int
	BitsPerComponent int
	iccProfile       *icc.Profile
	iccProfileErr    error
}

// ICCProfile returns an extracted ICC profile from this metadata.
//
// An error is returned if the ICC profile could not be correctly parsed.
//
// If no profile data was found, nil is returned without an error.
func (md *Data) ICCProfile() (*icc.Profile, error) {
	return md.iccProfile, md.iccProfileErr
}

func (md *Data) SetICCProfile(p *icc.Profile) {
	md.iccProfile = p
	md.iccProfileErr = nil
}

func (md *Data) SetICCProfileErr(err error) {
	md.iccProfile = nil
	md.iccProfileErr = err
}
