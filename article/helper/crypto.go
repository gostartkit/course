// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package helper

import (
	"crypto/sha256"
	"encoding/hex"

	"pkg.gostartkit.com/utils"
)

// CreateToken32 generate token
func CreateToken32() (string, error) {

	ud, err := utils.RandomString(32)

	if err != nil {
		return "", err
	}

	return ud, nil
}

// CreateToken64 generate token
func CreateToken64() (string, error) {

	ud, err := utils.RandomString(64)

	if err != nil {
		return "", err
	}

	return ud, nil
}

// Hash val
func Hash(val string) string {
	h := sha256.New()
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
}
