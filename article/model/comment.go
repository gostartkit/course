package model

import (
	"time"
)

// NewComment return *Comment
func NewComment() *Comment {

	comment := &Comment{}

	return comment
}

// Comment model
// @Entity tableName="comments"
type Comment struct {
	// @PrimaryKey
	ID uint64 `json:"id"`
	// @Ref Article.ID
	ArticleID uint64 `json:"articleID"`
	// @DataType mysql=varchar(127)
	CommentName string `json:"commentName"`
	// @Comment "-1 deleted 0 pendding 1 valid"
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// NewComments return *CommentCollection
func NewComments() *CommentCollection {

	comments := &CommentCollection{}

	return comments
}

// CommentCollection Comment list
type CommentCollection []Comment

// Len return len
func (o *CommentCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *CommentCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *CommentCollection) Less(i, j int) bool { return (*o)[i].ID < (*o)[j].ID }
