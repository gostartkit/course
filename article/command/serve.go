// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package command

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"app.gostartkit.com/go/article/config"
	"app.gostartkit.com/go/article/repository"
	"app.gostartkit.com/go/article/route"
	"pkg.gostartkit.com/cmd"
	"pkg.gostartkit.com/utils"
	"pkg.gostartkit.com/web"
)

var (
	cmdServe = &cmd.Command{
		Run:       runServe,
		UsageLine: "serve",
		Short:     "start web service",
		Long:      "start web service.",
	}
)

func runServe(cmd *cmd.Command, args []string) error {

	if err := config.Init(); err != nil {
		return fmt.Errorf("config.Init: %v", err)
	}

	if err := repository.Init(config.Database()); err != nil {
		return fmt.Errorf("repository.Init: %v", err)
	}

	defer repository.Close()

	app := web.CreateApplication()

	route.Init(app)

	network := config.Server().Network

	if network == "" {
		network = "unix"
	}

	addr := config.Server().Addr

	if addr == "" {
		if network == "tcp" {
			addr = "127.0.0.1:3000"
		} else {
			addr = "./log/unix.sock"
		}
	}

	appName := config.App().AppName

	if appName == "" {
		appName = fmt.Sprintf("%s.go", config.Key())
	}

	if network == "unix" {

		if utils.FileExist(addr) {
			// remove sock file if exists
			if err := os.Remove(addr); err != nil {
				return err
			}
		} else {
			// create dir if not exists
			logDir := path.Dir(addr)
			if err := os.MkdirAll(logDir, 0755); err != nil {
				return err
			}
		}
	}

	app.SetCORS(utils.Cors)

	log.Printf("%s(%d) %s:%s\n", appName, os.Getpid(), network, addr)

	err := app.ListenAndServe(network, addr, func(srv *http.Server) {
		srv.ReadTimeout = config.Server().ReadTimeout * time.Second
		srv.ReadHeaderTimeout = config.Server().ReadHeaderTimeout * time.Second
		srv.WriteTimeout = config.Server().WriteTimeout * time.Second
		srv.IdleTimeout = config.Server().IdleTimeout * time.Second
	})

	if err != nil {
		return fmt.Errorf("serve: %v", err)
	}

	return nil
}
