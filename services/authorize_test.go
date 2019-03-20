package services

import (
	"int_ecosys/models"
	"testing"
)

func TestStoreToken(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Store token failed: %v", err)
		}
	}()

	authInfo := &models.AuthInfo{
		AccessToken: "Access token",
		ExpiresIn:   1553095630,
	}
	_, err := StoreAuthInfo(authInfo)
	if err != nil {
		t.Errorf("Token is not stored with error %s\n", err)
	}
}
