// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package command

import (
	"fmt"

	"app.gostartkit.com/go/article/config"
	"pkg.gostartkit.com/cmd"
)

var (
	cmdVersion = &cmd.Command{
		Run:       runVersion,
		UsageLine: "version",
		Short:     "display version",
		Long:      "display version and build info.\n",
	}
)

func runVersion(cmd *cmd.Command, args []string) error {

	if Verbose() {
		fmt.Println(config.Key(), config.Version, GitRev(), OsArch())
		fmt.Println(config.StubName, config.StubVersion, config.StubGitRev)
	} else {
		fmt.Println(config.Version)
	}

	return nil
}
