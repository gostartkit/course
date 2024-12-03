// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package command

import (
	"errors"

	"app.gostartkit.com/go/article/config"
	"pkg.gostartkit.com/cmd"
)

var (
	cmdConfig = &cmd.Command{
		Run:       runConfig,
		UsageLine: "config [--tcp] [-v verbose] [-f force]",
		Short:     "create config file",
		Long:      "create config.json file at current directory.",
	}
)

func runConfig(cmd *cmd.Command, args []string) error {

	if len(args) != 0 {
		return errors.New("too many arguments given")
	}

	return config.WriteConfig(Tcp(), Verbose(), Force())
}
