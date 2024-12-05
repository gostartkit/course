package proxy

import (
	"app.gostartkit.com/go/article/repository"
	"pkg.gostartkit.com/utils"
)

var (
	dataRepository = repository.CreateDataRepository()
)

// GetAuthByAccessToken get auth by accessToken
func GetAuthByAccessToken(accessToken string) (*utils.Auth, error) {
	return dataRepository.GetAuthByAccessToken(accessToken)
}
