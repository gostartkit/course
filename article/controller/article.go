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
	_articleController     *ArticleController
	_onceArticleController sync.Once
)

// CreateArticleController return *ArticleController
func CreateArticleController() *ArticleController {

	_onceArticleController.Do(func() {
		_articleController = &ArticleController{}
	})

	return _articleController
}

// ArticleController struct
type ArticleController struct {
}

// Index get articles
func (r *ArticleController) Index(c *web.Ctx) (any, error) {

	filter := c.QueryFilter()
	orderBy := c.QueryOrderBy()
	page := c.QueryPage(_defaultPage)
	pageSize := c.QueryPageSize(_defaultPageSize)

	return proxy.GetArticles(filter, orderBy, page, pageSize)
}

// Detail get article
func (r *ArticleController) Detail(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.GetArticle(id)
}

// CreateID create article.ID
func (r *ArticleController) CreateID(c *web.Ctx) (any, error) {
	return proxy.CreateArticleID()
}

// Create create article
func (r *ArticleController) Create(c *web.Ctx) (any, error) {

	article := model.CreateArticle()
	defer article.Release()

	if err := c.TryParseBody(article); err != nil {
		return nil, err
	}

	if err := validator.CreateArticle(article); err != nil {
		return nil, err
	}

	if _, err := proxy.CreateArticle(article); err != nil {
		return nil, err
	}

	return article.ID, nil
}

// Update update article
func (r *ArticleController) Update(c *web.Ctx) (any, error) {

	var err error

	article := model.CreateArticle()
	defer article.Release()

	if err = c.TryParseBody(article); err != nil {
		return nil, err
	}

	if article.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateArticle(article); err != nil {
		return nil, err
	}

	return proxy.UpdateArticle(article)
}

// Patch update article
func (r *ArticleController) Patch(c *web.Ctx) (any, error) {

	attrs := c.HeaderAttrs()

	if attrs == nil {
		return nil, validator.ErrAttrsHeaderRequired
	}

	var err error

	article := model.CreateArticle()
	defer article.Release()

	if err = c.TryParseBody(article); err != nil {
		return nil, err
	}

	if article.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.PatchArticle(article, attrs...); err != nil {
		return nil, err
	}

	return proxy.PatchArticle(article, attrs...)
}

// UpdateStatus update article.Status
func (r *ArticleController) UpdateStatus(c *web.Ctx) (any, error) {

	var err error

	article := model.CreateArticle()
	defer article.Release()

	if err = c.TryParseBody(article); err != nil {
		return nil, err
	}

	if article.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateArticleStatus(article); err != nil {
		return nil, err
	}

	return proxy.UpdateArticleStatus(article)
}

// Destroy delete article
func (r *ArticleController) Destroy(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.DestroyArticleSoft(id)
}

// Tags return *model.TagCollection, error
func (r *ArticleController) Tags(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	filter := c.QueryFilter()
	orderBy := c.QueryOrderBy()
	page := c.QueryPage(_defaultPage)
	pageSize := c.QueryPageSize(_defaultPageSize)

	return proxy.GetTagsByArticleID(id, filter, orderBy, page, pageSize)
}

// LinkTags return rowsAffected int64, error
func (r *ArticleController) LinkTags(c *web.Ctx) (any, error) {

	var (
		tagID []uint64
	)

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := c.TryParseBody(&tagID); err != nil {
		return nil, err
	}

	return proxy.LinkArticleTags(id, tagID...)
}

// UnLinkTags return rowsAffected int64, error
func (r *ArticleController) UnLinkTags(c *web.Ctx) (any, error) {

	var (
		tagID []uint64
	)

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := c.TryParseBody(&tagID); err != nil {
		return nil, err
	}

	return proxy.UnLinkArticleTags(id, tagID...)
}
