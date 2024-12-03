// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package config

import "app.gostartkit.com/go/article/helper"

// CreateAppConfig create app config
func CreateAppConfig() *AppConfig {
	cfg := &AppConfig{
		AppID:                      1,
		AppNum:                     4,
		AppName:                    "article",
		AppEnv:                     "local",
		AppKey:                     helper.RandString(32),
		PrivateKey:                 "",
		AppDebug:                   false,
		Domain:                     _domain,
		PublicDir:                  "public",
		StorageDir:                 "storage",
		ResourceDir:                "resource",
		TimeLocation:               "Asia/Shanghai",
		TimeLayout:                 "2006-01-02 15:04:05",
		TokenExpireDuration:        3600 * 2,
		RefreshTokenExpireDuration: 3600 * 24 * 365,
	}

	return cfg
}

// AppConfig struct
type AppConfig struct {
	AppID                      uint64 `json:"appID"`
	AppNum                     uint64 `json:"appNum"`
	AppName                    string `json:"appName"`
	AppEnv                     string `json:"appEnv"`
	AppKey                     string `json:"appKey"`
	PrivateKey                 string `json:"privateKey"`
	AppDebug                   bool   `json:"appDebug"`
	Domain                     string `json:"domain"`
	PublicDir                  string `json:"publicDir"`
	StorageDir                 string `json:"storageDir"`
	ResourceDir                string `json:"resourceDir"`
	TimeLocation               string `json:"timeLocation"`
	TimeLayout                 string `json:"timeLayout"`
	TokenExpireDuration        uint   `json:"tokenExpireDuration"`
	RefreshTokenExpireDuration uint   `json:"refreshTokenExpireDuration"`
}
