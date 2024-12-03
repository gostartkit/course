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

func commentRoute(app *web.Application, prefix string) {

	c := controller.CreateCommentController()

	app.Get(prefix+"/comment/", middleware.Chain(c.Index, config.Read))
	app.Get(prefix+"/comment/:id", middleware.Chain(c.Detail, config.Read))
	app.Post(prefix+"/apply/comment/id/", middleware.Chain(c.CreateID, config.Write))
	app.Post(prefix+"/comment/", middleware.Chain(c.Create, config.Write))
	app.Put(prefix+"/comment/:id", middleware.Chain(c.Update, config.Write))
	app.Patch(prefix+"/comment/:id", middleware.Chain(c.Patch, config.Write))
	app.Patch(prefix+"/comment/:id/status/", middleware.Chain(c.UpdateStatus, config.Write))
	app.Delete(prefix+"/comment/:id", middleware.Chain(c.Destroy, config.Write))
}
