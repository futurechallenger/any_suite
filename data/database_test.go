package data

import "testing"

func TestNewIntEcoDB(t *testing.T) {
	if _, err := NewIntEcoDB(); err != nil {
		t.Error("Open db error")
	}
}
