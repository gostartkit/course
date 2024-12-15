package model

// NewArticleTag return *ArticleTag
func NewArticleTag() *ArticleTag {

	articleTag := &ArticleTag{}

	return articleTag
}

// ArticleTag model
// @Entity tableName="article_tag"
// @Ignore
type ArticleTag struct {
	// @PrimaryKey
	// @Ref Article.ID
	ArticleID uint64 `json:"articleID"`
	// @PrimaryKey
	// @Ref Tag.ID
	TagID uint64 `json:"tagID"`
}

// NewArticleTags return *ArticleTagCollection
func NewArticleTags() *ArticleTagCollection {

	articleTags := &ArticleTagCollection{}

	return articleTags
}

// ArticleTagCollection ArticleTag list
type ArticleTagCollection []ArticleTag

// Len return len
func (o *ArticleTagCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *ArticleTagCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *ArticleTagCollection) Less(i, j int) bool { return (*o)[i].ArticleID < (*o)[j].ArticleID }
