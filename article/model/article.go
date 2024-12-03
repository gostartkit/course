package model

import (
	"time"
)

// NewArticle return *Article
func NewArticle() *Article {

	article := &Article{}

	return article
}

// Article model
// @Entity tableName="articles"
type Article struct {
	// @PrimaryKey
	ID uint64 `json:"id"`
	// @Ref Category.ID
	CategoryID uint64 `json:"categoryID"`
	// @DataType mysql=varchar(127)
	ArticleName string `json:"articleName"`
	// @Comment "-1 deleted 0 pendding 1 valid"
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// NewArticles return *ArticleCollection
func NewArticles() *ArticleCollection {

	articles := &ArticleCollection{}

	return articles
}

// ArticleCollection Article list
type ArticleCollection []Article

// Len return len
func (o *ArticleCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *ArticleCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *ArticleCollection) Less(i, j int) bool { return (*o)[i].ID < (*o)[j].ID }
