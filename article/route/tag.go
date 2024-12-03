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

func tagRoute(app *web.Application, prefix string) {

	c := controller.CreateTagController()

	app.Get(prefix+"/tag/", middleware.Chain(c.Index, config.Read))
	app.Get(prefix+"/tag/:id", middleware.Chain(c.Detail, config.Read))
	app.Post(prefix+"/apply/tag/id/", middleware.Chain(c.CreateID, config.Write))
	app.Post(prefix+"/tag/", middleware.Chain(c.Create, config.Write))
	app.Put(prefix+"/tag/:id", middleware.Chain(c.Update, config.Write))
	app.Patch(prefix+"/tag/:id", middleware.Chain(c.Patch, config.Write))
	app.Patch(prefix+"/tag/:id/status/", middleware.Chain(c.UpdateStatus, config.Write))
	app.Delete(prefix+"/tag/:id", middleware.Chain(c.Destroy, config.Write))
	app.Get(prefix+"/tag/:id/article/", middleware.Chain(c.Articles, config.Read))
	app.Post(prefix+"/tag/:id/article/", middleware.Chain(c.LinkArticles, config.Write))
	app.Delete(prefix+"/tag/:id/article/", middleware.Chain(c.UnLinkArticles, config.Write))
}
