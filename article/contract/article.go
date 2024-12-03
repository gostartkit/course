// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package contract

import "app.gostartkit.com/go/article/model"

// IArticle interface
type IArticle interface {
	// CreateArticleID return article.ID error
	CreateArticleID() (uint64, error)
	// GetArticles return *model.ArticleCollection, error
	GetArticles(filter string, orderBy string, page int, pageSize int) (*model.ArticleCollection, error)
	// GetArticle return *model.Article, error
	GetArticle(id uint64) (*model.Article, error)
	// CreateArticle return int64, error
	// Attributes: ID uint64, CategoryID uint64, ArticleName string, Status int
	CreateArticle(article *model.Article) (int64, error)
	// UpdateArticle return int64, error
	// Attributes: CategoryID uint64, ArticleName string, Status int
	UpdateArticle(article *model.Article) (int64, error)
	// PatchArticle return int64, error
	// Attributes: CategoryID uint64, ArticleName string, Status int
	PatchArticle(article *model.Article, attrsName ...string) (int64, error)
	// UpdateArticleStatus return int64, error
	// Attributes: Status int
	UpdateArticleStatus(article *model.Article) (int64, error)
	// DestroyArticle return int64, error
	DestroyArticle(id uint64) (int64, error)
	// DestroyArticle return int64, error
	DestroyArticleSoft(id uint64) (int64, error)
	// GetArticlesByTagID return *model.ArticleCollection, error
	GetArticlesByTagID(tagID uint64, filter string, orderBy string, page int, pageSize int) (*model.ArticleCollection, error)
	// GetArticleByTagID return *model.Article, error
	GetArticleByTagID(articleID uint64, tagID uint64) (*model.Article, error)
	// LinkArticleTags return rowsAffected int64, error
	LinkArticleTags(articleID uint64, tagID ...uint64) (int64, error)
	// UnLinkArticleTags return rowsAffected int64, error
	UnLinkArticleTags(articleID uint64, tagID ...uint64) (int64, error)
}
