package models

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
