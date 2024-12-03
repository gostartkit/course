// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package route

import (
	"app.gostartkit.com/go/article/config"
	"app.gostartkit.com/go/article/controller"
	"app.gostartkit.com/go/article/middleware"
	"pkg.gostartkit.com/web"
)

func categoryRoute(app *web.Application, prefix string) {

	c := controller.CreateCategoryController()

	app.Get(prefix+"/category/", middleware.Chain(c.Index, config.Read))
	app.Get(prefix+"/category/:id", middleware.Chain(c.Detail, config.Read))
	app.Post(prefix+"/apply/category/id/", middleware.Chain(c.CreateID, config.Write))
	app.Post(prefix+"/category/", middleware.Chain(c.Create, config.Write))
	app.Put(prefix+"/category/:id", middleware.Chain(c.Update, config.Write))
	app.Patch(prefix+"/category/:id", middleware.Chain(c.Patch, config.Write))
	app.Patch(prefix+"/category/:id/status/", middleware.Chain(c.UpdateStatus, config.Write))
	app.Delete(prefix+"/category/:id", middleware.Chain(c.Destroy, config.Write))
}
