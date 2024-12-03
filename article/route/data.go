package route

import (
	"app.gostartkit.com/go/article/controller"
	"app.gostartkit.com/go/article/middleware"
	"pkg.gostartkit.com/web"
)

func dataRoute(app *web.Application, prefix string) {

	c := controller.CreateDataController()

	app.Get(prefix+"/config/rbac/", middleware.Chain(c.Rbac))
	app.Get(prefix+"/config/rbac/user/right/", middleware.Chain(c.RbacUserRight))
}
