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
	articleRepository = repository.CreateArticleRepository()
)

// CreateArticleID return article.ID error
func CreateArticleID() (uint64, error) {
	return articleRepository.CreateArticleID()
}

// GetArticles return *model.ArticleCollection, error
func GetArticles(filter string, orderBy string, page int, pageSize int) (*model.ArticleCollection, error) {
	return articleRepository.GetArticles(filter, orderBy, page, pageSize)
}

// GetArticle return *model.Article, error
func GetArticle(id uint64) (*model.Article, error) {
	return articleRepository.GetArticle(id)
}

// CreateArticle return int64, error
// Attributes: ID uint64, CategoryID uint64, ArticleName string, Status int
func CreateArticle(article *model.Article) (int64, error) {
	return articleRepository.CreateArticle(article)
}

// UpdateArticle return int64, error
// Attributes: CategoryID uint64, ArticleName string, Status int
func UpdateArticle(article *model.Article) (int64, error) {
	return articleRepository.UpdateArticle(article)
}

// PatchArticle return int64, error
// Attributes: CategoryID uint64, ArticleName string, Status int
func PatchArticle(article *model.Article, attrsName ...string) (int64, error) {
	return articleRepository.PatchArticle(article, attrsName...)
}

// UpdateArticleStatus return int64, error
// Attributes: Status int
func UpdateArticleStatus(article *model.Article) (int64, error) {
	return articleRepository.UpdateArticleStatus(article)
}

// DestroyArticle return int64, error
func DestroyArticle(id uint64) (int64, error) {
	return articleRepository.DestroyArticle(id)
}

// DestroyArticle return int64, error
func DestroyArticleSoft(id uint64) (int64, error) {
	return articleRepository.DestroyArticleSoft(id)
}

// GetArticlesByTagID return *model.ArticleCollection, error
func GetArticlesByTagID(tagID uint64, filter string, orderBy string, page int, pageSize int) (*model.ArticleCollection, error) {
	return articleRepository.GetArticlesByTagID(tagID, filter, orderBy, page, pageSize)
}

// GetArticleByTagID return *model.Article, error
func GetArticleByTagID(articleID uint64, tagID uint64) (*model.Article, error) {
	return articleRepository.GetArticleByTagID(articleID, tagID)
}

// LinkArticleTags return rowsAffected int64, error
func LinkArticleTags(articleID uint64, tagID ...uint64) (int64, error) {
	return articleRepository.LinkArticleTags(articleID, tagID...)
}

// UnLinkArticleTags return rowsAffected int64, error
func UnLinkArticleTags(articleID uint64, tagID ...uint64) (int64, error) {
	return articleRepository.UnLinkArticleTags(articleID, tagID...)
}
