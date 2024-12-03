package controller

import (
	"sync"

	"app.gostartkit.com/go/article/config"
	"pkg.gostartkit.com/web"
)

var (
	_dataController     *DataController
	_onceDataController sync.Once
)

// CreateDataController return *DataController
func CreateDataController() *DataController {

	_onceDataController.Do(func() {
		_dataController = &DataController{}
	})

	return _dataController
}

// DataController struct
type DataController struct {
}

// Index get user login form
func (r *DataController) Index(c *web.Ctx) (any, error) {
	return "", nil
}

// Rbac get rbac
func (r *DataController) Rbac(c *web.Ctx) (any, error) {
	return config.Rbac(), nil
}

// RbacUserRight get current user right
func (r *DataController) RbacUserRight(c *web.Ctx) (any, error) {
	return config.Rbac().Keys(c.UserRight()), nil
}
