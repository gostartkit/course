// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package validator

import (
	"app.gostartkit.com/go/article/model"
)

// CreateComment validate create comment
func CreateComment(comment *model.Comment) error {

	if comment.CommentName == "" {
		return createRequiredError("commentName")
	}

	return nil
}

// UpdateComment validate update comment
func UpdateComment(comment *model.Comment) error {

	if comment.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}

// PatchComment validate update comment part
func PatchComment(comment *model.Comment, attrsName ...string) error {

	if comment.ID == 0 {
		return createRequiredError("id")
	}

	if len(attrsName) == 0 {
		return createRequiredError("attrs")
	}

	return nil
}

// UpdateCommentStatus validate update comment status
func UpdateCommentStatus(comment *model.Comment) error {

	if comment.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}
