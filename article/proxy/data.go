package proxy

import (
	"app.gostartkit.com/go/article/model"
	"app.gostartkit.com/go/article/repository"
)

var (
	dataRepository = repository.CreateDataRepository()
)

// GetAuthByAccessToken get auth by accessToken
func GetAuthByAccessToken(accessToken string) (*model.Auth, error) {
	return dataRepository.GetAuthByAccessToken(accessToken)
}
