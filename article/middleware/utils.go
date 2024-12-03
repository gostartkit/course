// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package middleware

import (
	"pkg.gostartkit.com/web"
)

func setContentType(c *web.Ctx) {

	ct := c.Accept()

	switch ct {
	case "application/json", "application/x-gob", "application/octet-stream", "application/xml":
		c.SetContentType(ct)
	default:
		c.SetContentType("application/json")
	}
}
