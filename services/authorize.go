package services

import (
	"fmt"
	"int_ecosys/data"
	"int_ecosys/models"
)

// GetAuthInfo get auth token
func GetAuthInfo() {
}

// StoreAuthInfo stores auth info
func StoreAuthInfo(autoInfo *models.AuthInfo) (int64, error) {
	database := data.NewIntEcoDB()
	if database == nil {
		return -1, fmt.Errorf("IntEcoDB initialized error")
	}

	rowID, err := database.StoreToken(autoInfo)
	if err != nil {
		return 0, err
	}

	return rowID, nil
}
