// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package config

import (
	"errors"
	"sort"
	"sync"
	"time"

	"app.gostartkit.com/go/article/helper"
	"pkg.gostartkit.com/utils"
)

const (
	_configFile = "config.json"
	_key        = "article"

	_domain = "gostartkit.com"

	// DatabaseCluster
	_driver    = "mysql"
	_host      = "127.0.0.1"
	_port      = 3306
	_database  = "article"
	_username  = "article"
	_charset   = "utf8"
	_collation = "utf8_general_ci"
)

var (
	_timeLocation *time.Location
	_webConfig    *webConfig
	_once         sync.Once
)

// Init config
func Init() error {

	var err error

	_once.Do(func() {
		_webConfig, err = readConfig()

		if err == nil {
			_timeLocation, err = time.LoadLocation(_webConfig.App.TimeLocation)
		}
	})

	return err
}

// App get AppConfig
func App() *AppConfig {
	return _webConfig.App
}

// Server get ServerConfig
func Server() *ServerConfig {
	return _webConfig.Server
}

// Database get DatabaseClusterConfig
func Database() *DatabaseCluster {
	return _webConfig.Database
}

// Auth get AuthConfig
func Auth() *AuthConfig {
	return _webConfig.Auth
}

// Rbac get RbacConfig
func Rbac() *RightCollection {
	return _webConfig.rbac
}

// Key get Key
func Key() string {
	return _key
}

// TimeLocation get time location
func TimeLocation() *time.Location {
	return _timeLocation
}

// TimeLayout get time layout
func TimeLayout() string {

	if _webConfig.App.TimeLayout == "" {
		return "2006-01-02 15:04:05"
	}

	return _webConfig.App.TimeLayout
}

// AuthUrl get auth url
func AuthUrl() string {

	addr := Auth().Addr

	l := len(addr)

	if l == 1 {
		return addr[0]
	}

	return addr[helper.RandMax(l)]
}

// webConfig struct
type webConfig struct {
	App      *AppConfig       `json:"app"`
	Server   *ServerConfig    `json:"server"`
	Database *DatabaseCluster `json:"database"`
	Auth     *AuthConfig      `json:"auth"`

	rbac *RightCollection
}

// WriteConfig create new config.json at $configDir
func WriteConfig(tcp bool, verbose bool, force bool) error {

	cfg := &webConfig{}

	cfg.App = CreateAppConfig()

	cfg.Server = CreateServerConfig(tcp)

	cfg.Database = CreateDatabaseClusterConfig()

	cfg.Auth = CreateAuthConfig(_domain)

	cfg.rbac = CreateRbacConfig()

	if err := utils.WriteJSON(_configFile, cfg, force); err != nil {
		return err
	}

	return nil
}

// readConfig read $configDir/config.json
func readConfig() (*webConfig, error) {

	c := &webConfig{}

	if utils.FileExist(_configFile) {

		if err := utils.ReadJSON(_configFile, c); err != nil {
			return nil, err
		}
	}

	if c.App == nil {
		c.App = CreateAppConfig()
	}

	if c.Server == nil {
		c.Server = CreateServerConfig(false)
	}

	if c.Database == nil {
		c.Database = CreateDatabaseClusterConfig()
	}

	if c.Database.Write == nil {
		return nil, errors.New("config.Database.Write is nil")
	}

	if c.Database.Read == nil {
		return nil, errors.New("config.Database.Read is nil")
	}

	if len(*c.Database.Read) == 0 {
		return nil, errors.New("config.Database.Read len is 0")
	}

	if c.Auth == nil {
		c.Auth = CreateAuthConfig(c.App.Domain)
	}

	if len(c.Auth.Addr) == 0 {
		return nil, errors.New("config.Auth.Addr len is 0")
	}

	if c.rbac == nil {
		c.rbac = CreateRbacConfig()
	}

	sort.Sort(c.rbac)

	return c, nil
}
