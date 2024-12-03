// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package main

import (
	"app.gostartkit.com/go/article/command"
	_ "pkg.gostartkit.com/mysql"
)

var (
	osarch string
	gitrev string
)

func main() {
	command.Run(osarch, gitrev)
}
