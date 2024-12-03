// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package proxy

import (
	"app.gostartkit.com/go/article/model"
	"app.gostartkit.com/go/article/repository"
)

var (
	categoryRepository = repository.CreateCategoryRepository()
)

// CreateCategoryID return category.ID error
func CreateCategoryID() (uint64, error) {
	return categoryRepository.CreateCategoryID()
}

// GetCategories return *model.CategoryCollection, error
func GetCategories(filter string, orderBy string, page int, pageSize int) (*model.CategoryCollection, error) {
	return categoryRepository.GetCategories(filter, orderBy, page, pageSize)
}

// GetCategory return *model.Category, error
func GetCategory(id uint64) (*model.Category, error) {
	return categoryRepository.GetCategory(id)
}

// CreateCategory return int64, error
// Attributes: ID uint64, Ref uint64, CategoryName string, Status int
func CreateCategory(category *model.Category) (int64, error) {
	return categoryRepository.CreateCategory(category)
}

// UpdateCategory return int64, error
// Attributes: Ref uint64, CategoryName string, Status int
func UpdateCategory(category *model.Category) (int64, error) {
	return categoryRepository.UpdateCategory(category)
}

// PatchCategory return int64, error
// Attributes: Ref uint64, CategoryName string, Status int
func PatchCategory(category *model.Category, attrsName ...string) (int64, error) {
	return categoryRepository.PatchCategory(category, attrsName...)
}

// UpdateCategoryStatus return int64, error
// Attributes: Status int
func UpdateCategoryStatus(category *model.Category) (int64, error) {
	return categoryRepository.UpdateCategoryStatus(category)
}

// DestroyCategory return int64, error
func DestroyCategory(id uint64) (int64, error) {
	return categoryRepository.DestroyCategory(id)
}

// DestroyCategory return int64, error
func DestroyCategorySoft(id uint64) (int64, error) {
	return categoryRepository.DestroyCategorySoft(id)
}
