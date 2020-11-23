package dao

import "gotour/blog-service/internal/model"

func (dao *Dao) CreateArticleTag(articleTag *model.ArticleTag) error  {
	return articleTag.CreateArticleTag(dao.engine)
}