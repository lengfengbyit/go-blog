package service

import (
	"gotour/blog-service/internal/dao"
	"gotour/blog-service/internal/model"
	"gotour/blog-service/pkg/app"
	"gotour/blog-service/pkg/convert"
	"gotour/blog-service/pkg/errcode"
	"strings"
)

type CountArticleRequest struct {
	Title string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ListArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"coverImageUrl" binding:"max=255"`
	Content       string `form:"content"`
	TagIds        string `form:"tagIds" binding:"required"`
	model.Model
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CoverImageUrl string `form:"coverImageUrl" binding:"max=255"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (ser *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return ser.dao.CountArticle(param.Title, param.State)
}

func (ser *Service) GetArticle(id uint32) (*model.Article, error) {
	return ser.dao.GetArticle(id)
}

func (ser *Service) GetArticleList(param *ListArticleRequest, paper *app.Pager) (
	articleList []*model.Article, err error) {
	return ser.dao.ListArticle(param.Title, param.State, paper.Page, paper.PageSize)
}

func (ser *Service) CreateArticle(param *CreateArticleRequest) error {
	// 检测标签ID是否存在
	tagIds := strings.Split(param.TagIds, ",")
	if !ser.dao.CheckTagExists(tagIds) {
		return errcode.InvalidParams
	}

	return ser.dao.Transaction(func(dao *dao.Dao) error {
		article, err := dao.CreateArticle(
			param.Title,
			param.Desc,
			param.Content,
			param.CoverImageUrl,
			param.CreatedBy,
			param.State,
		)

		if err != nil {
			return err
		}

		// 保存 文章 标签关系
		for _, tagId := range tagIds {
			artTag := &model.ArticleTag{
				ArticleId: article.ID,
				TagId:     convert.StrTo(tagId).MustUInt32(),
			}
			err = dao.CreateArticleTag(artTag)
			if err != nil {
				return err
			}
		}
		return err
	})
}

func (ser *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return ser.dao.UpdateArticle(param.ID, param.Title, param.State, "fym")
}

func (ser *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return ser.dao.DeleteArticle(param.ID)
}
