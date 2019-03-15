package config

import (
	"os"
	"testing"
)

func TestENV(t *testing.T) {
	clientID := os.Getenv("INTECO_DEV_CLIENT_ID")
	clientSecret := os.Getenv("INTECO_DEV_CLIENT_SECRET")

	if clientID == "" {
		t.Error("Can not get env var client ID")
	}

	if clientSecret == "" {
		t.Error("Cannot get env var client secret")
	}
}
