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
	tagRepository = repository.CreateTagRepository()
)

// CreateTagID return tag.ID error
func CreateTagID() (uint64, error) {
	return tagRepository.CreateTagID()
}

// GetTags return *model.TagCollection, error
func GetTags(filter string, orderBy string, page int, pageSize int) (*model.TagCollection, error) {
	return tagRepository.GetTags(filter, orderBy, page, pageSize)
}

// GetTag return *model.Tag, error
func GetTag(id uint64) (*model.Tag, error) {
	return tagRepository.GetTag(id)
}

// CreateTag return int64, error
// Attributes: ID uint64, TagName string, Status int
func CreateTag(tag *model.Tag) (int64, error) {
	return tagRepository.CreateTag(tag)
}

// UpdateTag return int64, error
// Attributes: TagName string, Status int
func UpdateTag(tag *model.Tag) (int64, error) {
	return tagRepository.UpdateTag(tag)
}

// PatchTag return int64, error
// Attributes: TagName string, Status int
func PatchTag(tag *model.Tag, attrsName ...string) (int64, error) {
	return tagRepository.PatchTag(tag, attrsName...)
}

// UpdateTagStatus return int64, error
// Attributes: Status int
func UpdateTagStatus(tag *model.Tag) (int64, error) {
	return tagRepository.UpdateTagStatus(tag)
}

// DestroyTag return int64, error
func DestroyTag(id uint64) (int64, error) {
	return tagRepository.DestroyTag(id)
}

// DestroyTag return int64, error
func DestroyTagSoft(id uint64) (int64, error) {
	return tagRepository.DestroyTagSoft(id)
}

// GetTagsByArticleID return *model.TagCollection, error
func GetTagsByArticleID(articleID uint64, filter string, orderBy string, page int, pageSize int) (*model.TagCollection, error) {
	return tagRepository.GetTagsByArticleID(articleID, filter, orderBy, page, pageSize)
}

// GetTagByArticleID return *model.Tag, error
func GetTagByArticleID(tagID uint64, articleID uint64) (*model.Tag, error) {
	return tagRepository.GetTagByArticleID(tagID, articleID)
}

// LinkTagArticles return rowsAffected int64, error
func LinkTagArticles(tagID uint64, articleID ...uint64) (int64, error) {
	return tagRepository.LinkTagArticles(tagID, articleID...)
}

// UnLinkTagArticles return rowsAffected int64, error
func UnLinkTagArticles(tagID uint64, articleID ...uint64) (int64, error) {
	return tagRepository.UnLinkTagArticles(tagID, articleID...)
}
