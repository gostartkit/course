// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package contract

import "app.gostartkit.com/go/article/model"

// ITag interface
type ITag interface {
	// CreateTagID return tag.ID error
	CreateTagID() (uint64, error)
	// GetTags return *model.TagCollection, error
	GetTags(filter string, orderBy string, page int, pageSize int) (*model.TagCollection, error)
	// GetTag return *model.Tag, error
	GetTag(id uint64) (*model.Tag, error)
	// CreateTag return int64, error
	// Attributes: ID uint64, TagName string, Status int
	CreateTag(tag *model.Tag) (int64, error)
	// UpdateTag return int64, error
	// Attributes: TagName string, Status int
	UpdateTag(tag *model.Tag) (int64, error)
	// PatchTag return int64, error
	// Attributes: TagName string, Status int
	PatchTag(tag *model.Tag, attrsName ...string) (int64, error)
	// UpdateTagStatus return int64, error
	// Attributes: Status int
	UpdateTagStatus(tag *model.Tag) (int64, error)
	// DestroyTag return int64, error
	DestroyTag(id uint64) (int64, error)
	// DestroyTag return int64, error
	DestroyTagSoft(id uint64) (int64, error)
	// GetTagsByArticleID return *model.TagCollection, error
	GetTagsByArticleID(articleID uint64, filter string, orderBy string, page int, pageSize int) (*model.TagCollection, error)
	// GetTagByArticleID return *model.Tag, error
	GetTagByArticleID(tagID uint64, articleID uint64) (*model.Tag, error)
	// LinkTagArticles return rowsAffected int64, error
	LinkTagArticles(tagID uint64, articleID ...uint64) (int64, error)
	// UnLinkTagArticles return rowsAffected int64, error
	UnLinkTagArticles(tagID uint64, articleID ...uint64) (int64, error)
}
