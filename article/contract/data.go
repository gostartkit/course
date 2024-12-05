package contract

import "pkg.gostartkit.com/utils"

// DataRepository interface
type DataRepository interface {
	// GetAuthByAccessToken get auth by accessToken
	GetAuthByAccessToken(accessToken string) (*utils.Auth, error)
}
