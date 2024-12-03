// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package contract

import "app.gostartkit.com/go/article/model"

// ICategory interface
type ICategory interface {
	// CreateCategoryID return category.ID error
	CreateCategoryID() (uint64, error)
	// GetCategories return *model.CategoryCollection, error
	GetCategories(filter string, orderBy string, page int, pageSize int) (*model.CategoryCollection, error)
	// GetCategory return *model.Category, error
	GetCategory(id uint64) (*model.Category, error)
	// CreateCategory return int64, error
	// Attributes: ID uint64, Ref uint64, CategoryName string, Status int
	CreateCategory(category *model.Category) (int64, error)
	// UpdateCategory return int64, error
	// Attributes: Ref uint64, CategoryName string, Status int
	UpdateCategory(category *model.Category) (int64, error)
	// PatchCategory return int64, error
	// Attributes: Ref uint64, CategoryName string, Status int
	PatchCategory(category *model.Category, attrsName ...string) (int64, error)
	// UpdateCategoryStatus return int64, error
	// Attributes: Status int
	UpdateCategoryStatus(category *model.Category) (int64, error)
	// DestroyCategory return int64, error
	DestroyCategory(id uint64) (int64, error)
	// DestroyCategory return int64, error
	DestroyCategorySoft(id uint64) (int64, error)
}
