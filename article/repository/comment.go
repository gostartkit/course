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
	_commentRepository     contract.IComment
	_onceCommentRepository sync.Once
)

// CreateCommentRepository return contract.IComment
func CreateCommentRepository() contract.IComment {

	_onceCommentRepository.Do(func() {
		_commentRepository = &CommentRepository{}
	})

	return _commentRepository
}

// CommentRepository struct
type CommentRepository struct {
	mu        sync.Mutex
	commentID uint64
}

// CreateCommentID return comment.ID error
func (r *CommentRepository) CreateCommentID() (uint64, error) {
	r.mu.Lock()
	if r.commentID == 0 {
		var err error
		r.commentID, err = max("comments", "id", config.App().AppID, config.App().AppNum)
		if err != nil {
			r.mu.Unlock()
			return 0, err
		}
		if r.commentID == 0 {
			r.commentID = config.App().AppID - config.App().AppNum
		}
	}
	r.mu.Unlock()
	commentID := atomic.AddUint64(&r.commentID, config.App().AppNum)
	return commentID, nil
}

// GetComments return *model.CommentCollection, error
func (r *CommentRepository) GetComments(filter string, orderBy string, page int, pageSize int) (*model.CommentCollection, error) {

	var sqlx strings.Builder
	var args []any

	sqlx.WriteString("SELECT `id`, `article_id`, `comment_name`, `status`, `created_at`, `updated_at` ")
	sqlx.WriteString("FROM `comments` ")
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

	comments := model.CreateComments(pageSize)

	for rows.Next() {

		comment := model.CreateComment()

		err := rows.Scan(&comment.ID, &comment.ArticleID, &comment.CommentName, &comment.Status, &comment.CreatedAt, &comment.UpdatedAt)

		if err != nil {
			comment.Release()
			comments.Release()
			return nil, err
		}

		*comments = append(*comments, *comment)
	}

	return comments, rows.Err()
}

// GetComment return *model.Comment, error
func (r *CommentRepository) GetComment(id uint64) (*model.Comment, error) {

	sqlx := "SELECT `id`, `article_id`, `comment_name`, `status`, `created_at`, `updated_at` " +
		"FROM `comments` " +
		"WHERE `id` = ? AND `status` >= 0"

	row := queryRow(sqlx, id)

	comment := model.CreateComment()

	err := row.Scan(&comment.ID, &comment.ArticleID, &comment.CommentName, &comment.Status, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		comment.Release()
		if err == sql.ErrNoRows {
			return nil, web.ErrNotFound
		}
		return nil, err
	}

	return comment, nil
}

// CreateComment return int64, error
// Attributes: ID uint64, ArticleID uint64, CommentName string, Status int
func (r *CommentRepository) CreateComment(comment *model.Comment) (int64, error) {

	sqlx := "INSERT INTO `comments` " +
		"(`id`, `article_id`, `comment_name`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?, ?)"

	var err error

	if comment.ID == 0 {

		comment.ID, err = r.CreateCommentID()

		if err != nil {
			return 0, err
		}
	}

	comment.CreatedAt = now()

	result, err := exec(sqlx, comment.ID, comment.ArticleID, comment.CommentName, comment.Status, comment.CreatedAt)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateComment return int64, error
// Attributes: ArticleID uint64, CommentName string, Status int
func (r *CommentRepository) UpdateComment(comment *model.Comment) (int64, error) {

	sqlx := "UPDATE `comments` " +
		"SET `article_id` = ?, `comment_name` = ?, `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	comment.UpdatedAt = now()

	result, err := exec(sqlx, comment.ArticleID, comment.CommentName, comment.Status, comment.UpdatedAt, comment.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// PatchComment return int64, error
// Attributes: ArticleID uint64, CommentName string, Status int
func (r *CommentRepository) PatchComment(comment *model.Comment, attrsName ...string) (int64, error) {

	var sqlx strings.Builder
	var args []any

	rv := reflect.Indirect(reflect.ValueOf(comment))

	sqlx.WriteString("UPDATE `comments` SET ")

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

	comment.UpdatedAt = now()

	args = append(args, comment.UpdatedAt)

	sqlx.WriteString(" WHERE `id` = ?")

	args = append(args, comment.ID)

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

// UpdateCommentStatus return int64, error
// Attributes: Status int
func (r *CommentRepository) UpdateCommentStatus(comment *model.Comment) (int64, error) {

	sqlx := "UPDATE `comments` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	comment.UpdatedAt = now()

	result, err := exec(sqlx, comment.Status, comment.UpdatedAt, comment.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyComment return int64, error
func (r *CommentRepository) DestroyComment(id uint64) (int64, error) {

	sqlx := "DELETE FROM `comments` WHERE `id` = ?"

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

// DestroyComment return int64, error
func (r *CommentRepository) DestroyCommentSoft(id uint64) (int64, error) {

	sqlx := "UPDATE `comments` " +
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
func (r *CommentRepository) tryParse(key string, val string) (string, any, error) {

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
func (r *CommentRepository) tryParseKey(key string) (string, string, string, error) {

	switch key {
	case "id", "ID":
		return "id", "ID", "uint64", nil
	case "articleID", "ArticleID":
		return "article_id", "ArticleID", "uint64", nil
	case "commentName", "CommentName":
		return "comment_name", "CommentName", "string", nil
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
