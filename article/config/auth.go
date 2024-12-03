// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package config

import "fmt"

// CreateAuthConfig return *AuthConfig
func CreateAuthConfig(domain string) *AuthConfig {

	cfg := &AuthConfig{
		Addr: []string{
			fmt.Sprintf("https://auth.%s/authorize/%s/", domain, Key()),
		},
	}

	return cfg
}

// AuthConfig struct
type AuthConfig struct {
	Addr []string `json:"addr"`
}
