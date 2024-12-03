// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package config

import "app.gostartkit.com/go/article/helper"

// CreateDatabaseClusterConfig create database cluster config
func CreateDatabaseClusterConfig() *DatabaseCluster {
	cfg := &DatabaseCluster{
		Driver:    _driver,
		Database:  _database,
		Username:  _username,
		Password:  helper.RandString(16),
		Charset:   _charset,
		Collation: _collation,
	}

	cfg.Write = &DatabaseHost{
		Host: _host,
		Port: _port,
	}

	cfg.Read = &[]DatabaseHost{
		{
			Host: _host,
			Port: _port,
		},
		{
			Host: _host,
			Port: _port,
		},
	}

	return cfg
}

// DatabaseCluster struct
type DatabaseCluster struct {
	Driver    string          `json:"driver"`
	Database  string          `json:"database"`
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Charset   string          `json:"charset"`
	Collation string          `json:"collation"`
	Write     *DatabaseHost   `json:"write"`
	Read      *[]DatabaseHost `json:"read"`
}

// DatabaseHost struct
type DatabaseHost struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
