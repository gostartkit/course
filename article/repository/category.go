// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"

	"app.gostartkit.com/go/article/config"
	"app.gostartkit.com/go/article/contract"
	"app.gostartkit.com/go/article/model"
	"pkg.gostartkit.com/utils"
	"pkg.gostartkit.com/web"
)

var (
	_categoryRepository     contract.ICategory
	_onceCategoryRepository sync.Once
)

// CreateCategoryRepository return contract.ICategory
func CreateCategoryRepository() contract.ICategory {

	_onceCategoryRepository.Do(func() {
		_categoryRepository = &CategoryRepository{}
	})

	return _categoryRepository
}

// CategoryRepository struct
type CategoryRepository struct {
	mu         sync.Mutex
	categoryID uint64
}

// CreateCategoryID return category.ID error
func (r *CategoryRepository) CreateCategoryID() (uint64, error) {
	r.mu.Lock()
	if r.categoryID == 0 {
		var err error
		r.categoryID, err = max("categories", "id", config.App().AppID, config.App().AppNum)
		if err != nil {
			r.mu.Unlock()
			return 0, err
		}
		if r.categoryID == 0 {
			r.categoryID = config.App().AppID - config.App().AppNum
		}
	}
	r.mu.Unlock()
	categoryID := atomic.AddUint64(&r.categoryID, config.App().AppNum)
	return categoryID, nil
}

// GetCategories return *model.CategoryCollection, error
func (r *CategoryRepository) GetCategories(filter string, orderBy string, page int, pageSize int) (*model.CategoryCollection, error) {

	var sqlx strings.Builder
	var args []any

	sqlx.WriteString("SELECT `id`, `ref`, `category_name`, `status`, `created_at`, `updated_at` ")
	sqlx.WriteString("FROM `categories` ")
	sqlx.WriteString("WHERE `status` >= 0 ")

	if filter != "" {
		sqlx.WriteString("AND ")
		if err := utils.SqlFilter(filter, &sqlx, &args, "", r.tryParse); err != nil {
			return nil, err
		}
		sqlx.WriteString(" ")
	}

	if orderBy != "" {
		sqlx.WriteString("ORDER BY ")
		if err := utils.SqlOrderBy(orderBy, &sqlx, "", r.tryParseKey); err != nil {
			return nil, err
		}
		sqlx.WriteString(" ")
	}

	sqlx.WriteString("limit ? offset ?")

	if pageSize > _maxPageSize {
		pageSize = _maxPageSize
	} else if pageSize <= 0 {
		pageSize = _pageSize
	}

	offset := 0

	if page > 1 {
		offset = (page - 1) * pageSize
	}

	args = append(args, pageSize, offset)

	rows, err := query(sqlx.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := model.CreateCategories(pageSize)

	for rows.Next() {

		category := model.CreateCategory()

		err := rows.Scan(&category.ID, &category.Ref, &category.CategoryName, &category.Status, &category.CreatedAt, &category.UpdatedAt)

		if err != nil {
			category.Release()
			categories.Release()
			return nil, err
		}

		*categories = append(*categories, *category)
	}

	return categories, rows.Err()
}

// GetCategory return *model.Category, error
func (r *CategoryRepository) GetCategory(id uint64) (*model.Category, error) {

	sqlx := "SELECT `id`, `ref`, `category_name`, `status`, `created_at`, `updated_at` " +
		"FROM `categories` " +
		"WHERE `id` = ? AND `status` >= 0"

	row := queryRow(sqlx, id)

	category := model.CreateCategory()

	err := row.Scan(&category.ID, &category.Ref, &category.CategoryName, &category.Status, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		category.Release()
		if err == sql.ErrNoRows {
			return nil, web.ErrNotFound
		}
		return nil, err
	}

	return category, nil
}

// CreateCategory return int64, error
// Attributes: ID uint64, Ref uint64, CategoryName string, Status int
func (r *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {

	sqlx := "INSERT INTO `categories` " +
		"(`id`, `ref`, `category_name`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?, ?)"

	var err error

	if category.ID == 0 {

		category.ID, err = r.CreateCategoryID()

		if err != nil {
			return 0, err
		}
	}

	category.CreatedAt = now()

	result, err := exec(sqlx, category.ID, category.Ref, category.CategoryName, category.Status, category.CreatedAt)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateCategory return int64, error
// Attributes: Ref uint64, CategoryName string, Status int
func (r *CategoryRepository) UpdateCategory(category *model.Category) (int64, error) {

	sqlx := "UPDATE `categories` " +
		"SET `ref` = ?, `category_name` = ?, `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	category.UpdatedAt = now()

	result, err := exec(sqlx, category.Ref, category.CategoryName, category.Status, category.UpdatedAt, category.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// PatchCategory return int64, error
// Attributes: Ref uint64, CategoryName string, Status int
func (r *CategoryRepository) PatchCategory(category *model.Category, attrsName ...string) (int64, error) {

	var sqlx strings.Builder
	var args []any

	rv := reflect.Indirect(reflect.ValueOf(category))

	sqlx.WriteString("UPDATE `categories` SET ")

	for i, n := range attrsName {

		columnName, attributeName, _, err := r.tryParseKey(n)

		if err != nil {
			return 0, err
		}

		if i > 0 {
			sqlx.WriteString(", ")
		}

		fmt.Fprintf(&sqlx, "`%s` = ?", columnName)

		v := rv.FieldByName(attributeName).Interface()

		args = append(args, v)
	}

	sqlx.WriteString(", `updated_at` = ?")

	category.UpdatedAt = now()

	args = append(args, category.UpdatedAt)

	sqlx.WriteString(" WHERE `id` = ?")

	args = append(args, category.ID)

	result, err := exec(sqlx.String(), args...)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateCategoryStatus return int64, error
// Attributes: Status int
func (r *CategoryRepository) UpdateCategoryStatus(category *model.Category) (int64, error) {

	sqlx := "UPDATE `categories` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	category.UpdatedAt = now()

	result, err := exec(sqlx, category.Status, category.UpdatedAt, category.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyCategory return int64, error
func (r *CategoryRepository) DestroyCategory(id uint64) (int64, error) {

	sqlx := "DELETE FROM `categories` WHERE `id` = ?"

	result, err := exec(sqlx, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyCategory return int64, error
func (r *CategoryRepository) DestroyCategorySoft(id uint64) (int64, error) {

	sqlx := "UPDATE `categories` " +
		"SET `status` = -ABS(`status`) " +
		"WHERE `id` = ?"

	result, err := exec(sqlx, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// tryParse return columnName, attributeValue, error
func (r *CategoryRepository) tryParse(key string, val string) (string, any, error) {

	columnName, _, attributeType, err := r.tryParseKey(key)

	if err != nil {
		return "", nil, err
	}

	v, err := utils.TryParse(val, attributeType)

	if err != nil {
		return "", nil, err
	}

	return columnName, v, nil
}

// tryParseKey return columnName, attributeName, attributeType, error
func (r *CategoryRepository) tryParseKey(key string) (string, string, string, error) {

	switch key {
	case "id", "ID":
		return "id", "ID", "uint64", nil
	case "ref", "Ref":
		return "ref", "Ref", "uint64", nil
	case "categoryName", "CategoryName":
		return "category_name", "CategoryName", "string", nil
	case "status", "Status":
		return "status", "Status", "int", nil
	case "createdAt", "CreatedAt":
		return "created_at", "CreatedAt", "*time.Time", nil
	case "updatedAt", "UpdatedAt":
		return "updated_at", "UpdatedAt", "*time.Time", nil
	default:
		return "", "", "", web.ErrInvalid
	}
}
