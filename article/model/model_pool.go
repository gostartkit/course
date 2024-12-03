// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package model

import "sync"

var (
	_articlePool = sync.Pool{
		New: func() any {
			return NewArticle()
		}}

	_articleTagPool = sync.Pool{
		New: func() any {
			return NewArticleTag()
		}}

	_categoryPool = sync.Pool{
		New: func() any {
			return NewCategory()
		}}

	_commentPool = sync.Pool{
		New: func() any {
			return NewComment()
		}}

	_tagPool = sync.Pool{
		New: func() any {
			return NewTag()
		}}
)

// CreateArticle return *Article
func CreateArticle() *Article {

	article := _articlePool.Get().(*Article)

	return article
}

func (o *Article) initial() {
	o.ID = 0
	o.CategoryID = 0
	o.ArticleName = ""
	o.Status = 0
	o.CreatedAt = nil
	o.UpdatedAt = nil
}

func (o *Article) Release() {
	o.initial()
	_articlePool.Put(o)
}

// CreateArticles return *ArticleCollection
func CreateArticles(pageSize int) *ArticleCollection {

	articles := make(ArticleCollection, 0, pageSize)

	return &articles
}

func (o *ArticleCollection) Release() {
	for i := 0; i < len(*o); i++ {
		(*o)[i].Release()
	}
}

// CreateArticleTag return *ArticleTag
func CreateArticleTag() *ArticleTag {

	articleTag := _articleTagPool.Get().(*ArticleTag)

	return articleTag
}

func (o *ArticleTag) initial() {
	o.ArticleID = 0
	o.TagID = 0
}

func (o *ArticleTag) Release() {
	o.initial()
	_articleTagPool.Put(o)
}

// CreateArticleTags return *ArticleTagCollection
func CreateArticleTags(pageSize int) *ArticleTagCollection {

	articleTags := make(ArticleTagCollection, 0, pageSize)

	return &articleTags
}

func (o *ArticleTagCollection) Release() {
	for i := 0; i < len(*o); i++ {
		(*o)[i].Release()
	}
}

// CreateCategory return *Category
func CreateCategory() *Category {

	category := _categoryPool.Get().(*Category)

	return category
}

func (o *Category) initial() {
	o.ID = 0
	o.Ref = 0
	o.CategoryName = ""
	o.Status = 0
	o.CreatedAt = nil
	o.UpdatedAt = nil
}

func (o *Category) Release() {
	o.initial()
	_categoryPool.Put(o)
}

// CreateCategories return *CategoryCollection
func CreateCategories(pageSize int) *CategoryCollection {

	categories := make(CategoryCollection, 0, pageSize)

	return &categories
}

func (o *CategoryCollection) Release() {
	for i := 0; i < len(*o); i++ {
		(*o)[i].Release()
	}
}

// CreateComment return *Comment
func CreateComment() *Comment {

	comment := _commentPool.Get().(*Comment)

	return comment
}

func (o *Comment) initial() {
	o.ID = 0
	o.ArticleID = 0
	o.CommentName = ""
	o.Status = 0
	o.CreatedAt = nil
	o.UpdatedAt = nil
}

func (o *Comment) Release() {
	o.initial()
	_commentPool.Put(o)
}

// CreateComments return *CommentCollection
func CreateComments(pageSize int) *CommentCollection {

	comments := make(CommentCollection, 0, pageSize)

	return &comments
}

func (o *CommentCollection) Release() {
	for i := 0; i < len(*o); i++ {
		(*o)[i].Release()
	}
}

// CreateTag return *Tag
func CreateTag() *Tag {

	tag := _tagPool.Get().(*Tag)

	return tag
}

func (o *Tag) initial() {
	o.ID = 0
	o.TagName = ""
	o.Status = 0
	o.CreatedAt = nil
	o.UpdatedAt = nil
}

func (o *Tag) Release() {
	o.initial()
	_tagPool.Put(o)
}

// CreateTags return *TagCollection
func CreateTags(pageSize int) *TagCollection {

	tags := make(TagCollection, 0, pageSize)

	return &tags
}

func (o *TagCollection) Release() {
	for i := 0; i < len(*o); i++ {
		(*o)[i].Release()
	}
}
