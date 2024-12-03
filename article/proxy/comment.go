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
	commentRepository = repository.CreateCommentRepository()
)

// CreateCommentID return comment.ID error
func CreateCommentID() (uint64, error) {
	return commentRepository.CreateCommentID()
}

// GetComments return *model.CommentCollection, error
func GetComments(filter string, orderBy string, page int, pageSize int) (*model.CommentCollection, error) {
	return commentRepository.GetComments(filter, orderBy, page, pageSize)
}

// GetComment return *model.Comment, error
func GetComment(id uint64) (*model.Comment, error) {
	return commentRepository.GetComment(id)
}

// CreateComment return int64, error
// Attributes: ID uint64, ArticleID uint64, CommentName string, Status int
func CreateComment(comment *model.Comment) (int64, error) {
	return commentRepository.CreateComment(comment)
}

// UpdateComment return int64, error
// Attributes: ArticleID uint64, CommentName string, Status int
func UpdateComment(comment *model.Comment) (int64, error) {
	return commentRepository.UpdateComment(comment)
}

// PatchComment return int64, error
// Attributes: ArticleID uint64, CommentName string, Status int
func PatchComment(comment *model.Comment, attrsName ...string) (int64, error) {
	return commentRepository.PatchComment(comment, attrsName...)
}

// UpdateCommentStatus return int64, error
// Attributes: Status int
func UpdateCommentStatus(comment *model.Comment) (int64, error) {
	return commentRepository.UpdateCommentStatus(comment)
}

// DestroyComment return int64, error
func DestroyComment(id uint64) (int64, error) {
	return commentRepository.DestroyComment(id)
}

// DestroyComment return int64, error
func DestroyCommentSoft(id uint64) (int64, error) {
	return commentRepository.DestroyCommentSoft(id)
}
