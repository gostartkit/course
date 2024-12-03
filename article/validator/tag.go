// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package validator

import (
	"app.gostartkit.com/go/article/model"
)

// CreateTag validate create tag
func CreateTag(tag *model.Tag) error {

	if tag.TagName == "" {
		return createRequiredError("tagName")
	}

	return nil
}

// UpdateTag validate update tag
func UpdateTag(tag *model.Tag) error {

	if tag.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}

// PatchTag validate update tag part
func PatchTag(tag *model.Tag, attrsName ...string) error {

	if tag.ID == 0 {
		return createRequiredError("id")
	}

	if len(attrsName) == 0 {
		return createRequiredError("attrs")
	}

	return nil
}

// UpdateTagStatus validate update tag status
func UpdateTagStatus(tag *model.Tag) error {

	if tag.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}
