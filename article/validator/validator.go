// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package validator

import (
	"errors"
	"fmt"
)

var (
	// ErrAttrsHeaderRequired attrs header required
	ErrAttrsHeaderRequired = errors.New("attrs header required")
)

// createRequiredError create required error
func createRequiredError(name string) error {
	return createValidationError(name, "required")
}

// createInvalidError create invalid error
func createInvalidError(name string) error {
	return createValidationError(name, "invalid")
}

// createValidationError create validation error with name and message
func createValidationError(name string, message string) error {
	return fmt.Errorf("%s %s", name, message)
}
