package services

import (
	"any_suite/data"
	"any_suite/models"
	"fmt"
)

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
	db := data.NewAppDB()
	if db == nil {
		return false, fmt.Errorf("Initialize db failed")
	}

	return false, nil
}
