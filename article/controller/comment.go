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
	_commentController     *CommentController
	_onceCommentController sync.Once
)

// CreateCommentController return *CommentController
func CreateCommentController() *CommentController {

	_onceCommentController.Do(func() {
		_commentController = &CommentController{}
	})

	return _commentController
}

// CommentController struct
type CommentController struct {
}

// Index get comments
func (r *CommentController) Index(c *web.Ctx) (any, error) {

	filter := c.QueryFilter()
	orderBy := c.QueryOrderBy()
	page := c.QueryPage(_defaultPage)
	pageSize := c.QueryPageSize(_defaultPageSize)

	return proxy.GetComments(filter, orderBy, page, pageSize)
}

// Detail get comment
func (r *CommentController) Detail(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.GetComment(id)
}

// CreateID create comment.ID
func (r *CommentController) CreateID(c *web.Ctx) (any, error) {
	return proxy.CreateCommentID()
}

// Create create comment
func (r *CommentController) Create(c *web.Ctx) (any, error) {

	comment := model.CreateComment()
	defer comment.Release()

	if err := c.TryParseBody(comment); err != nil {
		return nil, err
	}

	if err := validator.CreateComment(comment); err != nil {
		return nil, err
	}

	if _, err := proxy.CreateComment(comment); err != nil {
		return nil, err
	}

	return comment.ID, nil
}

// Update update comment
func (r *CommentController) Update(c *web.Ctx) (any, error) {

	var err error

	comment := model.CreateComment()
	defer comment.Release()

	if err = c.TryParseBody(comment); err != nil {
		return nil, err
	}

	if comment.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateComment(comment); err != nil {
		return nil, err
	}

	return proxy.UpdateComment(comment)
}

// Patch update comment
func (r *CommentController) Patch(c *web.Ctx) (any, error) {

	attrs := c.HeaderAttrs()

	if attrs == nil {
		return nil, validator.ErrAttrsHeaderRequired
	}

	var err error

	comment := model.CreateComment()
	defer comment.Release()

	if err = c.TryParseBody(comment); err != nil {
		return nil, err
	}

	if comment.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.PatchComment(comment, attrs...); err != nil {
		return nil, err
	}

	return proxy.PatchComment(comment, attrs...)
}

// UpdateStatus update comment.Status
func (r *CommentController) UpdateStatus(c *web.Ctx) (any, error) {

	var err error

	comment := model.CreateComment()
	defer comment.Release()

	if err = c.TryParseBody(comment); err != nil {
		return nil, err
	}

	if comment.ID, err = c.ParamUint64("id"); err != nil {
		return nil, err
	}

	if err = validator.UpdateCommentStatus(comment); err != nil {
		return nil, err
	}

	return proxy.UpdateCommentStatus(comment)
}

// Destroy delete comment
func (r *CommentController) Destroy(c *web.Ctx) (any, error) {

	id, err := c.ParamUint64("id")

	if err != nil {
		return nil, err
	}

	if err := utils.Uint64("id", id); err != nil {
		return nil, err
	}

	return proxy.DestroyCommentSoft(id)
}
