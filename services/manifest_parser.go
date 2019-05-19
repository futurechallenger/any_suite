package services

// ManifestParser parses `appmanifest.json`
type ManifestParser struct {
}

// NewManifestParser returns the pointer of a new instance of `ManifestParser`.
func NewManifestParser(manifestpath string) *ManifestParser {
	return &ManifestParser{}
}
