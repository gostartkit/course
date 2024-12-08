# article

a base app for RESTful API

## Gegetting started

### Build from source

```bash
go build -ldflags "-s -w" -buildmode=exe -tags release -o bin/article
```

### Create config

```bash
bin/article config
```

**Note**: Please modify config.json as you need.

### Start service

```bash
bin/article serve
```

**Note**: Default config is use network: `unix`, you can change `network` to `tcp` and `addr` to `127.0.0.1:5000` for test.

### Create models

```bash
gsk model category article comment tag articleTag
```

### Update Models

```go
// Category model
// @Entity tableName="categories"
type Category struct {
	// @PrimaryKey
	ID uint64 `json:"id"`
	// @Comment "parent category id"
	Ref uint64 `json:"ref"`
	// @DataType mysql=varchar(127)
	CategoryName string `json:"categoryName"`
	// @Comment "-1 deleted 0 pendding 1 valid"
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
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

// ArticleTag model
// @Entity tableName="article_tags"
// @Ignore
type ArticleTag struct {
	// @PrimaryKey
	// @Ref Article.ID
	ArticleID uint64 `json:"articleID"`
	// @PrimaryKey
	// @Ref Tag.ID
	TagID uint64 `json:"tagID"`
}
```

### Create api code

```bash
gsk code
```

### Build

```bash
gsk build
```