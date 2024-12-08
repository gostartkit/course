// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package route

import (
	"sync"

	"pkg.gostartkit.com/web"
)

var (
	_once   sync.Once
	_prefix string
)

// Init route init
func Init(app *web.Application) {
	_once.Do(func() {
		dataRoute(app, _prefix)
		articleRoute(app, _prefix)
		categoryRoute(app, _prefix)
		commentRoute(app, _prefix)
		tagRoute(app, _prefix)
	})
}
