// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package contract

import "app.gostartkit.com/go/article/model"

// IComment interface
type IComment interface {
	// CreateCommentID return comment.ID error
	CreateCommentID() (uint64, error)
	// GetComments return *model.CommentCollection, error
	GetComments(filter string, orderBy string, page int, pageSize int) (*model.CommentCollection, error)
	// GetComment return *model.Comment, error
	GetComment(id uint64) (*model.Comment, error)
	// CreateComment return int64, error
	// Attributes: ID uint64, ArticleID uint64, CommentName string, Status int
	CreateComment(comment *model.Comment) (int64, error)
	// UpdateComment return int64, error
	// Attributes: ArticleID uint64, CommentName string, Status int
	UpdateComment(comment *model.Comment) (int64, error)
	// PatchComment return int64, error
	// Attributes: ArticleID uint64, CommentName string, Status int
	PatchComment(comment *model.Comment, attrsName ...string) (int64, error)
	// UpdateCommentStatus return int64, error
	// Attributes: Status int
	UpdateCommentStatus(comment *model.Comment) (int64, error)
	// DestroyComment return int64, error
	DestroyComment(id uint64) (int64, error)
	// DestroyComment return int64, error
	DestroyCommentSoft(id uint64) (int64, error)
}
