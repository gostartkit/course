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

func articleRoute(app *web.Application, prefix string) {

	c := controller.CreateArticleController()

	app.Get(prefix+"/article/", middleware.Chain(c.Index, config.Read))
	app.Get(prefix+"/article/:id", middleware.Chain(c.Detail, config.Read))
	app.Post(prefix+"/apply/article/id/", middleware.Chain(c.CreateID, config.Write))
	app.Post(prefix+"/article/", middleware.Chain(c.Create, config.Write))
	app.Put(prefix+"/article/:id", middleware.Chain(c.Update, config.Write))
	app.Patch(prefix+"/article/:id", middleware.Chain(c.Patch, config.Write))
	app.Patch(prefix+"/article/:id/status/", middleware.Chain(c.UpdateStatus, config.Write))
	app.Delete(prefix+"/article/:id", middleware.Chain(c.Destroy, config.Write))
	app.Get(prefix+"/article/:id/tag/", middleware.Chain(c.Tags, config.Read))
	app.Post(prefix+"/article/:id/tag/", middleware.Chain(c.LinkTags, config.Write))
	app.Delete(prefix+"/article/:id/tag/", middleware.Chain(c.UnLinkTags, config.Write))
}
