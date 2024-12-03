package model

import (
	"time"
)

// NewTag return *Tag
func NewTag() *Tag {

	tag := &Tag{}

	return tag
}

// Tag model
// @Entity tableName="tags"
type Tag struct {
	// @PrimaryKey
	ID uint64 `json:"id"`
	// @DataType mysql=varchar(127)
	TagName string `json:"tagName"`
	// @Comment "-1 deleted 0 pendding 1 valid"
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// NewTags return *TagCollection
func NewTags() *TagCollection {

	tags := &TagCollection{}

	return tags
}

// TagCollection Tag list
type TagCollection []Tag

// Len return len
func (o *TagCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *TagCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *TagCollection) Less(i, j int) bool { return (*o)[i].ID < (*o)[j].ID }
