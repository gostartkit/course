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
	_tagController     *TagController
	_onceTagController sync.Once
)

// CreateTagController return *TagController
func CreateTagController() *TagController {

	_onceTagController.Do(func() {
		_tagController = &TagController{}
	})

	return _tagController
}

// TagController struct
type TagController struct {
}

// Index get tags
func (r *TagController) Index(c *web.Ctx) (any, error) {

	filter := c.QueryFilter()
	orderBy := c.QueryOrderBy()
	page := c.QueryPage(_defaultPage)
	pageSize := c.QueryPageSize(_defaultPageSize)

	return proxy.GetTags(filter, orderBy, page, pageSize)
}

// Detail get tag
func (r *TagController) Detail(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.GetTag(id)
}

// CreateID create tag.ID
func (r *TagController) CreateID(c *web.Ctx) (any, error) {
	return proxy.CreateTagID()
}

// Create create tag
func (r *TagController) Create(c *web.Ctx) (any, error) {

	tag := model.CreateTag()
	defer tag.Release()

	if err := c.TryParseBody(tag); err != nil {
		return nil, err
	}

	if err := validator.CreateTag(tag); err != nil {
		return nil, err
	}

	if _, err := proxy.CreateTag(tag); err != nil {
		return nil, err
	}

	return tag.ID, nil
}

// Update update tag
func (r *TagController) Update(c *web.Ctx) (any, error) {

	var err error

	tag := model.CreateTag()
	defer tag.Release()

	if err = c.TryParseBody(tag); err != nil {
		return nil, err
	}

	if tag.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateTag(tag); err != nil {
		return nil, err
	}

	return proxy.UpdateTag(tag)
}

// Patch update tag
func (r *TagController) Patch(c *web.Ctx) (any, error) {

	attrs := c.HeaderAttrs()

	if attrs == nil {
		return nil, validator.ErrAttrsHeaderRequired
	}

	var err error

	tag := model.CreateTag()
	defer tag.Release()

	if err = c.TryParseBody(tag); err != nil {
		return nil, err
	}

	if tag.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.PatchTag(tag, attrs...); err != nil {
		return nil, err
	}

	return proxy.PatchTag(tag, attrs...)
}

// UpdateStatus update tag.Status
func (r *TagController) UpdateStatus(c *web.Ctx) (any, error) {

	var err error

	tag := model.CreateTag()
	defer tag.Release()

	if err = c.TryParseBody(tag); err != nil {
		return nil, err
	}

	if tag.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateTagStatus(tag); err != nil {
		return nil, err
	}

	return proxy.UpdateTagStatus(tag)
}

// Destroy delete tag
func (r *TagController) Destroy(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.DestroyTagSoft(id)
}

// Articles return *model.ArticleCollection, error
func (r *TagController) Articles(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	filter := c.QueryFilter()
	orderBy := c.QueryOrderBy()
	page := c.QueryPage(_defaultPage)
	pageSize := c.QueryPageSize(_defaultPageSize)

	return proxy.GetArticlesByTagID(id, filter, orderBy, page, pageSize)
}

// LinkArticles return rowsAffected int64, error
func (r *TagController) LinkArticles(c *web.Ctx) (any, error) {

	var (
		articleID []uint64
	)

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := c.TryParseBody(&articleID); err != nil {
		return nil, err
	}

	return proxy.LinkTagArticles(id, articleID...)
}

// UnLinkArticles return rowsAffected int64, error
func (r *TagController) UnLinkArticles(c *web.Ctx) (any, error) {

	var (
		articleID []uint64
	)

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := c.TryParseBody(&articleID); err != nil {
		return nil, err
	}

	return proxy.UnLinkTagArticles(id, articleID...)
}
