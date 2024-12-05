package repository

import (
	"sync"

	"app.gostartkit.com/go/article/config"
	"app.gostartkit.com/go/article/contract"
	"pkg.gostartkit.com/utils"
	"pkg.gostartkit.com/web"
)

var (
	_dataRepository     contract.DataRepository
	_onceDataRepository sync.Once
)

// CreateDataRepository return contract.AuthRepository
func CreateDataRepository() contract.DataRepository {

	_onceDataRepository.Do(func() {
		_dataRepository = &DataRepository{}
	})

	return _dataRepository
}

// DataRepository struct
type DataRepository struct {
}

// GetAuthByAccessToken get auth by accessToken
func (r *DataRepository) GetAuthByAccessToken(accessToken string) (*utils.Auth, error) {

	var err error

	auth := utils.CreateAuth()

	err = web.Get(config.AuthUrl(), accessToken, auth)

	if err != nil {
		auth.Release()
		return nil, err
	}

	return auth, nil
}
