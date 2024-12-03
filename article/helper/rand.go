// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package helper

import (
	"math/rand"

	"pkg.gostartkit.com/utils"
)

// RandString rand string
func RandString(len int) string {

	ud, err := utils.RandomString(len)

	if err != nil {
		return ""
	}

	return ud
}

// RandInt rand int between [min, max)
func RandInt(min int, max int) int {

	if min <= 0 || max <= 0 {
		return 0
	}

	if min >= max {
		return max
	}

	return rand.Intn(max-min) + min
}

// RandMax rand int between [0, max)
func RandMax(max int) int {

	if max <= 1 {
		return 0
	}

	return rand.Intn(max)
}
