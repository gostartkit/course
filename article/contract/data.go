package contract

import "app.gostartkit.com/go/article/model"

// DataRepository interface
type DataRepository interface {
	// GetAuthByAccessToken get auth by accessToken
	GetAuthByAccessToken(accessToken string) (*model.Auth, error)
}
