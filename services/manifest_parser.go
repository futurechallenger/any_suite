package services

import "any_suite/models"

// ManifestParser parses `appmanifest.json`
type ManifestParser struct {
	models.Manifest
}

// NewManifestParser returns the pointer of a new instance of `ManifestParser`.
func NewManifestParser(manifestpath string) *ManifestParser {
	return &ManifestParser{}
}
