package services

import "any_suite/models"

// ManifestParser parses `appmanifest.json`
type ManifestParser struct {
	mfst *models.Manifest
}

// NewManifestParser returns the pointer of a new instance of `ManifestParser`.
func NewManifestParser(manifest *models.Manifest) *ManifestParser {
	return &ManifestParser{mfst: manifest}
}

// Run starts to parse manifest
// Store parsed into storage, like `redis` or `mongodb`
func (mp *ManifestParser) Run() (bool, error) {
	return false, nil
}
