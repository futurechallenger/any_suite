package models

import "fmt"

// ParsedManifest presents `Manifest` instance and parsed info
type ParsedManifest struct {
	*Manifest
	Functions []string
}

// ParseManifest parse manifest instance
func ParseManifest(manifest *Manifest) (pm *ParsedManifest, err error) {
	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
		}
	}()

	if manifest == nil {
		return nil, fmt.Errorf("Manifest is invalid / empty")
	}

	pm = &ParsedManifest{manifest, make([]string, 10)}

	for _, v := range manifest.Triggers {
		if v == nil {
			continue
		}

		innerMap, ok := v.(map[string]string)
		if !ok {
			return nil, fmt.Errorf("Can not get event handler functions")
		}

		for _, fn := range innerMap {
			pm.Functions = append(pm.Functions, fn)
		}
	}

	return pm, nil
}
