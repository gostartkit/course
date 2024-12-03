// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package middleware

import (
	"app.gostartkit.com/go/article/config"
	"app.gostartkit.com/go/article/proxy"
	"pkg.gostartkit.com/utils"
	"pkg.gostartkit.com/web"
)

// Chain returns a web.Next function that chains middleware operations.
func Chain(next web.Next, vals ...int64) web.Next {

	return func(c *web.Ctx) (any, error) {

		origin := c.Origin()

		if origin != "" {
			if !utils.CheckCors(origin, config.App().Domain) {
				return nil, web.ErrCors
			}
			c.SetOrigin(origin)
		}

		before(c)

		err := bearerAuth(c, vals...)

		if err != nil {
			return nil, err
		}

		val, err := next(c)

		after(c)

		return val, err
	}
}

// Direct returns a web.Next function that directly executes the 'next' middleware.
func Direct(next web.Next, vals ...int64) web.Next {

	return func(c *web.Ctx) (any, error) {

		origin := c.Origin()

		if origin != "" {
			if !utils.CheckCors(origin, config.App().Domain) {
				return nil, web.ErrCors
			}
			c.SetOrigin(origin)
		}

		before(c)

		val, err := next(c)

		after(c)

		return val, err
	}
}

// before is a function that sets the 'Access-Control-Allow-Origin' header in the web context 'c'.
func before(c *web.Ctx) {
	setContentType(c)
}

// after is a function that performs post-processing tasks in the web context 'c'.
func after(c *web.Ctx) {

}

// bearerAuth is a function that performs bearer token authentication and authorization based on the provided access token and values.
func bearerAuth(c *web.Ctx, vals ...int64) error {

	accessToken := c.BearerToken()

	if accessToken == "" {
		return web.ErrUnauthorized
	}

	cat, err := proxy.GetAuthByAccessToken(accessToken)

	if err != nil {
		return err
	}

	c.Init(cat.UserID, cat.UserRight)

	if !utils.CheckVals(cat.UserRight, vals...) {
		cat.Release()
		return web.ErrForbidden
	}

	cat.Release()

	return nil
}
