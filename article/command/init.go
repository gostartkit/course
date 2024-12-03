// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package command

import (
	"errors"
	"fmt"

	"app.gostartkit.com/go/article/config"
	"app.gostartkit.com/go/article/repository"
	"pkg.gostartkit.com/cmd"
)

var (
	cmdInit = &cmd.Command{
		Run:       runInit,
		UsageLine: "init",
		Short:     "init system",
		Long:      `init system.`,
	}
)

func runInit(cmd *cmd.Command, args []string) error {

	if len(args) != 0 {
		return errors.New("too many arguments given")
	}

	if err := config.Init(); err != nil {
		return fmt.Errorf("config.Init: %v", err)
	}

	if err := repository.Init(config.Database()); err != nil {
		return fmt.Errorf("repository.Init: %v", err)
	}

	defer repository.Close()

	// -- code here

	return nil
}
