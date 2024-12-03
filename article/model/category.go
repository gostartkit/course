package model

import (
	"time"
)

// NewCategory return *Category
func NewCategory() *Category {

	category := &Category{}

	return category
}

// Category model
// @Entity tableName="categories"
type Category struct {
	// @PrimaryKey
	ID uint64 `json:"id"`
	// @DataType mysql=varchar(127)
	CategoryName string `json:"categoryName"`
	// @Comment "-1 deleted 0 pendding 1 valid"
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// NewCategories return *CategoryCollection
func NewCategories() *CategoryCollection {

	categories := &CategoryCollection{}

	return categories
}

// CategoryCollection Category list
type CategoryCollection []Category

// Len return len
func (o *CategoryCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *CategoryCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *CategoryCollection) Less(i, j int) bool { return (*o)[i].ID < (*o)[j].ID }
