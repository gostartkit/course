// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package command

import (
	"flag"

	"pkg.gostartkit.com/cmd"
)

var (
	_osarch  string
	_gitrev  string
	_verbose bool
	_force   bool
	_tcp     bool
)

// Run exec commands
func Run(osarch string, gitrev string) {

	_osarch = osarch
	_gitrev = gitrev

	cmd.SetFlags(func(f *flag.FlagSet) {
		f.BoolVar(&_verbose, "verbose", false, "verbose")
		f.BoolVar(&_verbose, "v", false, "verbose")
		f.BoolVar(&_force, "force", false, "force without warn")
		f.BoolVar(&_force, "f", false, "force without warn")

		f.BoolVar(&_tcp, "tcp", false, "use tcp")
	})
	cmd.AddCommands(cmdConfig, cmdInit, cmdServe, cmdVersion)
	cmd.Execute()
}

func OsArch() string {
	return _osarch
}

func GitRev() string {
	return _gitrev
}

func Verbose() bool {
	return _verbose
}

func Force() bool {
	return _force
}

func Tcp() bool {
	return _tcp
}
