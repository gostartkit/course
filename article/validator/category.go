// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package validator

import (
	"app.gostartkit.com/go/article/model"
)

// CreateCategory validate create category
func CreateCategory(category *model.Category) error {

	if category.CategoryName == "" {
		return createRequiredError("categoryName")
	}

	return nil
}

// UpdateCategory validate update category
func UpdateCategory(category *model.Category) error {

	if category.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}

// PatchCategory validate update category part
func PatchCategory(category *model.Category, attrsName ...string) error {

	if category.ID == 0 {
		return createRequiredError("id")
	}

	if len(attrsName) == 0 {
		return createRequiredError("attrs")
	}

	return nil
}

// UpdateCategoryStatus validate update category status
func UpdateCategoryStatus(category *model.Category) error {

	if category.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}
