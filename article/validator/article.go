// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package validator

import (
	"app.gostartkit.com/go/article/model"
)

// CreateArticle validate create article
func CreateArticle(article *model.Article) error {

	if article.ArticleName == "" {
		return createRequiredError("articleName")
	}

	return nil
}

// UpdateArticle validate update article
func UpdateArticle(article *model.Article) error {

	if article.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}

// PatchArticle validate update article part
func PatchArticle(article *model.Article, attrsName ...string) error {

	if article.ID == 0 {
		return createRequiredError("id")
	}

	if len(attrsName) == 0 {
		return createRequiredError("attrs")
	}

	return nil
}

// UpdateArticleStatus validate update article status
func UpdateArticleStatus(article *model.Article) error {

	if article.ID == 0 {
		return createRequiredError("id")
	}

	return nil
}
