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
	_tagRepository     contract.ITag
	_onceTagRepository sync.Once
)

// CreateTagRepository return contract.ITag
func CreateTagRepository() contract.ITag {

	_onceTagRepository.Do(func() {
		_tagRepository = &TagRepository{}
	})

	return _tagRepository
}

// TagRepository struct
type TagRepository struct {
	mu    sync.Mutex
	tagID uint64
}

// CreateTagID return tag.ID error
func (r *TagRepository) CreateTagID() (uint64, error) {
	r.mu.Lock()
	if r.tagID == 0 {
		var err error
		r.tagID, err = max("tags", "id", config.App().AppID, config.App().AppNum)
		if err != nil {
			r.mu.Unlock()
			return 0, err
		}
		if r.tagID == 0 {
			r.tagID = config.App().AppID - config.App().AppNum
		}
	}
	r.mu.Unlock()
	tagID := atomic.AddUint64(&r.tagID, config.App().AppNum)
	return tagID, nil
}

// GetTags return *model.TagCollection, error
func (r *TagRepository) GetTags(filter string, orderBy string, page int, pageSize int) (*model.TagCollection, error) {

	var sqlx strings.Builder
	var args []any

	sqlx.WriteString("SELECT `id`, `tag_name`, `status`, `created_at`, `updated_at` ")
	sqlx.WriteString("FROM `tags` ")
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

	tags := model.CreateTags(pageSize)

	for rows.Next() {

		tag := model.CreateTag()

		err := rows.Scan(&tag.ID, &tag.TagName, &tag.Status, &tag.CreatedAt, &tag.UpdatedAt)

		if err != nil {
			tag.Release()
			tags.Release()
			return nil, err
		}

		*tags = append(*tags, *tag)
	}

	return tags, rows.Err()
}

// GetTag return *model.Tag, error
func (r *TagRepository) GetTag(id uint64) (*model.Tag, error) {

	sqlx := "SELECT `id`, `tag_name`, `status`, `created_at`, `updated_at` " +
		"FROM `tags` " +
		"WHERE `id` = ? AND `status` >= 0"

	row := queryRow(sqlx, id)

	tag := model.CreateTag()

	err := row.Scan(&tag.ID, &tag.TagName, &tag.Status, &tag.CreatedAt, &tag.UpdatedAt)

	if err != nil {
		tag.Release()
		if err == sql.ErrNoRows {
			return nil, web.ErrNotFound
		}
		return nil, err
	}

	return tag, nil
}

// CreateTag return int64, error
// Attributes: ID uint64, TagName string, Status int
func (r *TagRepository) CreateTag(tag *model.Tag) (int64, error) {

	sqlx := "INSERT INTO `tags` " +
		"(`id`, `tag_name`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?)"

	var err error

	if tag.ID == 0 {

		tag.ID, err = r.CreateTagID()

		if err != nil {
			return 0, err
		}
	}

	tag.CreatedAt = now()

	result, err := exec(sqlx, tag.ID, tag.TagName, tag.Status, tag.CreatedAt)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateTag return int64, error
// Attributes: TagName string, Status int
func (r *TagRepository) UpdateTag(tag *model.Tag) (int64, error) {

	sqlx := "UPDATE `tags` " +
		"SET `tag_name` = ?, `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	tag.UpdatedAt = now()

	result, err := exec(sqlx, tag.TagName, tag.Status, tag.UpdatedAt, tag.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// PatchTag return int64, error
// Attributes: TagName string, Status int
func (r *TagRepository) PatchTag(tag *model.Tag, attrsName ...string) (int64, error) {

	var sqlx strings.Builder
	var args []any

	rv := reflect.Indirect(reflect.ValueOf(tag))

	sqlx.WriteString("UPDATE `tags` SET ")

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

	tag.UpdatedAt = now()

	args = append(args, tag.UpdatedAt)

	sqlx.WriteString(" WHERE `id` = ?")

	args = append(args, tag.ID)

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

// UpdateTagStatus return int64, error
// Attributes: Status int
func (r *TagRepository) UpdateTagStatus(tag *model.Tag) (int64, error) {

	sqlx := "UPDATE `tags` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	tag.UpdatedAt = now()

	result, err := exec(sqlx, tag.Status, tag.UpdatedAt, tag.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyTag return int64, error
func (r *TagRepository) DestroyTag(id uint64) (int64, error) {

	sqlx := "DELETE FROM `tags` WHERE `id` = ?"

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

// DestroyTag return int64, error
func (r *TagRepository) DestroyTagSoft(id uint64) (int64, error) {

	sqlx := "UPDATE `tags` " +
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

// GetTagsByArticleID return *model.TagCollection, error
func (r *TagRepository) GetTagsByArticleID(articleID uint64, filter string, orderBy string, page int, pageSize int) (*model.TagCollection, error) {

	var sqlx strings.Builder
	var args []any

	sqlx.WriteString("SELECT c.`id`, c.`tag_name`, c.`status`, c.`created_at`, c.`updated_at` ")
	sqlx.WriteString("FROM `tags` c ")
	sqlx.WriteString("INNER JOIN `article_tag` r on c.`id` = r.`tag_id` ")
	sqlx.WriteString("WHERE r.`article_id` = ? AND `status` >= 0 ")

	args = append(args, articleID)

	if filter != "" {
		sqlx.WriteString("AND ")
		if err := utils.SqlFilter(filter, &sqlx, &args, "c.", r.tryParse); err != nil {
			return nil, err
		}
		sqlx.WriteString(" ")
	}

	if orderBy != "" {
		sqlx.WriteString("ORDER BY ")
		if err := utils.SqlOrderBy(orderBy, &sqlx, "c.", r.tryParseKey); err != nil {
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

	tags := model.CreateTags(pageSize)

	for rows.Next() {

		tag := model.CreateTag()

		err := rows.Scan(&tag.ID, &tag.TagName, &tag.Status, &tag.CreatedAt, &tag.UpdatedAt)

		if err != nil {
			tag.Release()
			tags.Release()
			return nil, err
		}

		*tags = append(*tags, *tag)
	}

	return tags, rows.Err()
}

// GetTagByArticleID return *model.Tag, error
func (r *TagRepository) GetTagByArticleID(tagID uint64, articleID uint64) (*model.Tag, error) {

	sqlx := "SELECT c.`id`, c.`tag_name`, c.`status`, c.`created_at`, c.`updated_at` " +
		"FROM `tags` c " +
		"INNER JOIN `article_tag` r on c.`id` = r.`tag_id` " +
		"WHERE c.`id` = ? AND r.`article_id` = ? AND c.`status` >= 0 "

	row := queryRow(sqlx, tagID, articleID)

	tag := model.CreateTag()

	err := row.Scan(&tag.ID, &tag.TagName, &tag.Status, &tag.CreatedAt, &tag.UpdatedAt)

	if err != nil {
		tag.Release()
		if err == sql.ErrNoRows {
			return nil, web.ErrNotFound
		}
		return nil, err
	}

	return tag, nil
}

// LinkTagArticles return rowsAffected int64, error
func (r *TagRepository) LinkTagArticles(tagID uint64, articleID ...uint64) (int64, error) {

	sqlx := "INSERT INTO `article_tag` " +
		"(`article_id`, `tag_id`) " +
		"VALUES(?, ?)"

	tx, err := begin()

	if err != nil {
		return 0, err
	}

	stmt, err := txPrepare(tx, sqlx)

	if err != nil {
		rollback(tx)
		return 0, err
	}

	defer stmt.Close()

	var totalAffected int64 = 0

	for _, ref := range articleID {

		result, err := stmtExec(stmt, ref, tagID)

		if err != nil {
			rollback(tx)
			return 0, err
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			rollback(tx)
			return 0, err
		}

		totalAffected += rowsAffected
	}

	if err := commit(tx); err != nil {
		return 0, err
	}

	return totalAffected, nil
}

// UnLinkTagArticles return rowsAffected int64, error
func (r *TagRepository) UnLinkTagArticles(tagID uint64, articleID ...uint64) (int64, error) {

	sqlx := "DELETE FROM `article_tag` WHERE `article_id` = ? AND `tag_id` = ?"

	tx, err := begin()

	if err != nil {
		return 0, err
	}

	stmt, err := txPrepare(tx, sqlx)

	if err != nil {
		rollback(tx)
		return 0, err
	}

	defer stmt.Close()

	var totalAffected int64 = 0

	for _, ref := range articleID {

		result, err := stmtExec(stmt, ref, tagID)

		if err != nil {
			rollback(tx)
			return 0, err
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			rollback(tx)
			return 0, err
		}

		totalAffected += rowsAffected
	}

	if err := commit(tx); err != nil {
		return 0, err
	}

	return totalAffected, nil
}

// tryParse return columnName, attributeValue, error
func (r *TagRepository) tryParse(key string, val string) (string, any, error) {

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
func (r *TagRepository) tryParseKey(key string) (string, string, string, error) {

	switch key {
	case "id", "ID":
		return "id", "ID", "uint64", nil
	case "tagName", "TagName":
		return "tag_name", "TagName", "string", nil
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
