package icc

type Profile struct {
	Header   Header
	TagTable TagTable
}

func (p *Profile) Description() (string, error) {
	return p.TagTable.getProfileDescription(p.Header.Version)
}

func newProfile() *Profile {
	return &Profile{
		TagTable: emptyTagTable(),
	}
}
