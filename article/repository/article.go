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
	_articleRepository     contract.IArticle
	_onceArticleRepository sync.Once
)

// CreateArticleRepository return contract.IArticle
func CreateArticleRepository() contract.IArticle {

	_onceArticleRepository.Do(func() {
		_articleRepository = &ArticleRepository{}
	})

	return _articleRepository
}

// ArticleRepository struct
type ArticleRepository struct {
	mu        sync.Mutex
	articleID uint64
}

// CreateArticleID return article.ID error
func (r *ArticleRepository) CreateArticleID() (uint64, error) {
	r.mu.Lock()
	if r.articleID == 0 {
		var err error
		r.articleID, err = max("articles", "id", config.App().AppID, config.App().AppNum)
		if err != nil {
			r.mu.Unlock()
			return 0, err
		}
		if r.articleID == 0 {
			r.articleID = config.App().AppID - config.App().AppNum
		}
	}
	r.mu.Unlock()
	articleID := atomic.AddUint64(&r.articleID, config.App().AppNum)
	return articleID, nil
}

// GetArticles return *model.ArticleCollection, error
func (r *ArticleRepository) GetArticles(filter string, orderBy string, page int, pageSize int) (*model.ArticleCollection, error) {

	var sqlx strings.Builder
	var args []any

	sqlx.WriteString("SELECT `id`, `category_id`, `article_name`, `status`, `created_at`, `updated_at` ")
	sqlx.WriteString("FROM `articles` ")
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

	articles := model.CreateArticles(pageSize)

	for rows.Next() {

		article := model.CreateArticle()

		err := rows.Scan(&article.ID, &article.CategoryID, &article.ArticleName, &article.Status, &article.CreatedAt, &article.UpdatedAt)

		if err != nil {
			article.Release()
			articles.Release()
			return nil, err
		}

		*articles = append(*articles, *article)
	}

	return articles, rows.Err()
}

// GetArticle return *model.Article, error
func (r *ArticleRepository) GetArticle(id uint64) (*model.Article, error) {

	sqlx := "SELECT `id`, `category_id`, `article_name`, `status`, `created_at`, `updated_at` " +
		"FROM `articles` " +
		"WHERE `id` = ? AND `status` >= 0"

	row := queryRow(sqlx, id)

	article := model.CreateArticle()

	err := row.Scan(&article.ID, &article.CategoryID, &article.ArticleName, &article.Status, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		article.Release()
		if err == sql.ErrNoRows {
			return nil, web.ErrNotFound
		}
		return nil, err
	}

	return article, nil
}

// CreateArticle return int64, error
// Attributes: ID uint64, CategoryID uint64, ArticleName string, Status int
func (r *ArticleRepository) CreateArticle(article *model.Article) (int64, error) {

	sqlx := "INSERT INTO `articles` " +
		"(`id`, `category_id`, `article_name`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?, ?)"

	var err error

	if article.ID == 0 {

		article.ID, err = r.CreateArticleID()

		if err != nil {
			return 0, err
		}
	}

	article.CreatedAt = now()

	result, err := exec(sqlx, article.ID, article.CategoryID, article.ArticleName, article.Status, article.CreatedAt)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateArticle return int64, error
// Attributes: CategoryID uint64, ArticleName string, Status int
func (r *ArticleRepository) UpdateArticle(article *model.Article) (int64, error) {

	sqlx := "UPDATE `articles` " +
		"SET `category_id` = ?, `article_name` = ?, `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	article.UpdatedAt = now()

	result, err := exec(sqlx, article.CategoryID, article.ArticleName, article.Status, article.UpdatedAt, article.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// PatchArticle return int64, error
// Attributes: CategoryID uint64, ArticleName string, Status int
func (r *ArticleRepository) PatchArticle(article *model.Article, attrsName ...string) (int64, error) {

	var sqlx strings.Builder
	var args []any

	rv := reflect.Indirect(reflect.ValueOf(article))

	sqlx.WriteString("UPDATE `articles` SET ")

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

	article.UpdatedAt = now()

	args = append(args, article.UpdatedAt)

	sqlx.WriteString(" WHERE `id` = ?")

	args = append(args, article.ID)

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

// UpdateArticleStatus return int64, error
// Attributes: Status int
func (r *ArticleRepository) UpdateArticleStatus(article *model.Article) (int64, error) {

	sqlx := "UPDATE `articles` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ?"

	article.UpdatedAt = now()

	result, err := exec(sqlx, article.Status, article.UpdatedAt, article.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyArticle return int64, error
func (r *ArticleRepository) DestroyArticle(id uint64) (int64, error) {

	sqlx := "DELETE FROM `articles` WHERE `id` = ?"

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

// DestroyArticle return int64, error
func (r *ArticleRepository) DestroyArticleSoft(id uint64) (int64, error) {

	sqlx := "UPDATE `articles` " +
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

// GetArticlesByTagID return *model.ArticleCollection, error
func (r *ArticleRepository) GetArticlesByTagID(tagID uint64, filter string, orderBy string, page int, pageSize int) (*model.ArticleCollection, error) {

	var sqlx strings.Builder
	var args []any

	sqlx.WriteString("SELECT c.`id`, c.`category_id`, c.`article_name`, c.`status`, c.`created_at`, c.`updated_at` ")
	sqlx.WriteString("FROM `articles` c ")
	sqlx.WriteString("INNER JOIN `article_tag` r on c.`id` = r.`article_id` ")
	sqlx.WriteString("WHERE r.`tag_id` = ? AND `status` >= 0 ")

	args = append(args, tagID)

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

	articles := model.CreateArticles(pageSize)

	for rows.Next() {

		article := model.CreateArticle()

		err := rows.Scan(&article.ID, &article.CategoryID, &article.ArticleName, &article.Status, &article.CreatedAt, &article.UpdatedAt)

		if err != nil {
			article.Release()
			articles.Release()
			return nil, err
		}

		*articles = append(*articles, *article)
	}

	return articles, rows.Err()
}

// GetArticleByTagID return *model.Article, error
func (r *ArticleRepository) GetArticleByTagID(articleID uint64, tagID uint64) (*model.Article, error) {

	sqlx := "SELECT c.`id`, c.`category_id`, c.`article_name`, c.`status`, c.`created_at`, c.`updated_at` " +
		"FROM `articles` c " +
		"INNER JOIN `article_tag` r on c.`id` = r.`article_id` " +
		"WHERE c.`id` = ? AND r.`tag_id` = ? AND c.`status` >= 0 "

	row := queryRow(sqlx, articleID, tagID)

	article := model.CreateArticle()

	err := row.Scan(&article.ID, &article.CategoryID, &article.ArticleName, &article.Status, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		article.Release()
		if err == sql.ErrNoRows {
			return nil, web.ErrNotFound
		}
		return nil, err
	}

	return article, nil
}

// LinkArticleTags return rowsAffected int64, error
func (r *ArticleRepository) LinkArticleTags(articleID uint64, tagID ...uint64) (int64, error) {

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

	for _, ref := range tagID {

		result, err := stmtExec(stmt, articleID, ref)

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

// UnLinkArticleTags return rowsAffected int64, error
func (r *ArticleRepository) UnLinkArticleTags(articleID uint64, tagID ...uint64) (int64, error) {

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

	for _, ref := range tagID {

		result, err := stmtExec(stmt, articleID, ref)

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
func (r *ArticleRepository) tryParse(key string, val string) (string, any, error) {

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
func (r *ArticleRepository) tryParseKey(key string) (string, string, string, error) {

	switch key {
	case "id", "ID":
		return "id", "ID", "uint64", nil
	case "categoryID", "CategoryID":
		return "category_id", "CategoryID", "uint64", nil
	case "articleName", "ArticleName":
		return "article_name", "ArticleName", "string", nil
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
