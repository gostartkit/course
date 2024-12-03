// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package config

import (
	"time"
)

// CreateServerConfig create server config
func CreateServerConfig(tcp bool) *ServerConfig {

	var network string
	var addr string

	if tcp {
		network = "tcp"
		addr = "127.0.0.1:3000"
	} else {
		network = "unix"
		addr = "./log/unix.sock"
	}

	cfg := &ServerConfig{
		Network:           network,
		Addr:              addr,
		ReadTimeout:       32,
		ReadHeaderTimeout: 8,
		WriteTimeout:      32,
		IdleTimeout:       8,
	}
	return cfg
}

// ServerConfig struct
type ServerConfig struct {
	Network           string        `json:"network"`
	Addr              string        `json:"addr"`
	ReadTimeout       time.Duration `json:"readTimeout"`
	ReadHeaderTimeout time.Duration `json:"readHeaderTimeout"`
	WriteTimeout      time.Duration `json:"writeTimeout"`
	IdleTimeout       time.Duration `json:"idleTimeout"`
}
