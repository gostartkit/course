// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package controller

import (
	"sync"

	"app.gostartkit.com/go/article/model"
	"app.gostartkit.com/go/article/proxy"
	"app.gostartkit.com/go/article/validator"
	"pkg.gostartkit.com/utils"
	"pkg.gostartkit.com/web"
)

var (
	_categoryController     *CategoryController
	_onceCategoryController sync.Once
)

// CreateCategoryController return *CategoryController
func CreateCategoryController() *CategoryController {

	_onceCategoryController.Do(func() {
		_categoryController = &CategoryController{}
	})

	return _categoryController
}

// CategoryController struct
type CategoryController struct {
}

// Index get categories
func (r *CategoryController) Index(c *web.Ctx) (any, error) {

	filter := c.QueryFilter()
	orderBy := c.QueryOrderBy()
	page := c.QueryPage(_defaultPage)
	pageSize := c.QueryPageSize(_defaultPageSize)

	return proxy.GetCategories(filter, orderBy, page, pageSize)
}

// Detail get category
func (r *CategoryController) Detail(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.GetCategory(id)
}

// CreateID create category.ID
func (r *CategoryController) CreateID(c *web.Ctx) (any, error) {
	return proxy.CreateCategoryID()
}

// Create create category
func (r *CategoryController) Create(c *web.Ctx) (any, error) {

	category := model.CreateCategory()
	defer category.Release()

	if err := c.TryParseBody(category); err != nil {
		return nil, err
	}

	if err := validator.CreateCategory(category); err != nil {
		return nil, err
	}

	if _, err := proxy.CreateCategory(category); err != nil {
		return nil, err
	}

	return category.ID, nil
}

// Update update category
func (r *CategoryController) Update(c *web.Ctx) (any, error) {

	var err error

	category := model.CreateCategory()
	defer category.Release()

	if err = c.TryParseBody(category); err != nil {
		return nil, err
	}

	if category.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateCategory(category); err != nil {
		return nil, err
	}

	return proxy.UpdateCategory(category)
}

// Patch update category
func (r *CategoryController) Patch(c *web.Ctx) (any, error) {

	attrs := c.HeaderAttrs()

	if attrs == nil {
		return nil, validator.ErrAttrsHeaderRequired
	}

	var err error

	category := model.CreateCategory()
	defer category.Release()

	if err = c.TryParseBody(category); err != nil {
		return nil, err
	}

	if category.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.PatchCategory(category, attrs...); err != nil {
		return nil, err
	}

	return proxy.PatchCategory(category, attrs...)
}

// UpdateStatus update category.Status
func (r *CategoryController) UpdateStatus(c *web.Ctx) (any, error) {

	var err error

	category := model.CreateCategory()
	defer category.Release()

	if err = c.TryParseBody(category); err != nil {
		return nil, err
	}

	if category.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateCategoryStatus(category); err != nil {
		return nil, err
	}

	return proxy.UpdateCategoryStatus(category)
}

// Destroy delete category
func (r *CategoryController) Destroy(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.DestroyCategorySoft(id)
}
