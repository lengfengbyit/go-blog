package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"`
	TagId     uint32 `json:"tag_id"`
}

func (articleTag ArticleTag) TableName() string {
	return "blog_article_tag"
}

func (artTag *ArticleTag) CreateArticleTag(db *gorm.DB) error {
	return db.FirstOrCreate(artTag, &ArticleTag{ArticleId: artTag.ArticleId, TagId: artTag.TagId}).Error
}
