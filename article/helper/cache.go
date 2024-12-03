// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package helper

import "crypto/ecdsa"

var (
	_privateKey *ecdsa.PrivateKey
)

// PrivateKey cache privateKey
func PrivateKey(create func() (*ecdsa.PrivateKey, error)) (*ecdsa.PrivateKey, error) {

	if _privateKey == nil {

		var err error

		_privateKey, err = create()

		if err != nil {
			return nil, err
		}
	}

	return _privateKey, nil
}
