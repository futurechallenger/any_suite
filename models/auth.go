package models

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// AuthInfo auth information are here
type AuthInfo struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int32  `json:"expires_in"`
	RefreshTokenExpiresIn int32  `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	OwnerID               string `json:"owner_id"`
	EndPointID            string `json:"endpoint_id"`
}

type authDBInfo struct {
	AccessToken           string
	TokenType             sql.NullString
	RefreshToken          sql.NullString
	ExpiresIn             mysql.NullTime
	RefreshTokenExpiresIn mysql.NullTime
	Scope                 sql.NullString
	OwnerID               sql.NullString
	EndPointID            sql.NullString
}

// TODO: Iterate all fileds of a struct and get the value
// Then set values to db info instance
// https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go
func toDBAuthInfo(authInfo *AuthInfo) *authDBInfo {
	if authInfo == nil {
		return nil
	}

}
