package dao

import (
	"gotour/blog-service/internal/model"
	"gotour/blog-service/pkg/app"
	"time"
)

func (dao *Dao) CountArticle(title string, state uint8) (int, error) {
	article := &model.Article{Title: title, State: state}
	return article.Count(dao.engine)
}

func (dao *Dao) GetArticle(id uint32) (*model.Article, error) {
	article := &model.Article{
		Model: &model.Model{
			ID: id,
		},
	}
	return article.First(dao.engine)
}

func (dao *Dao) ListArticle(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := &model.Article{Title: title, State: state}
	pageOffSet := app.GetPageOffset(page, pageSize)
	return article.List(dao.engine, pageOffSet, pageSize)
}

func (dao *Dao) CreateArticle(title, desc, content, coverImageUrl, createdBy string, state uint8) (*model.Article, error) {
	article := &model.Article{
		Title:         title,
		Desc:          desc,
		CoverImageUrl: coverImageUrl,
		Content:       content,
		State:         state,
		Model: &model.Model{
			CreatedBy: createdBy,
			CreatedOn: uint32(time.Now().Unix()),
		},
	}

	err := article.Create(dao.engine)

	return article, err
}

func (dao *Dao) UpdateArticle(id uint32, title string, state uint8, modifiedBy string) error {
	article := &model.Article{
		Model: &model.Model{
			ID: id,
		},
	}

	var values = map[string]interface{}{
		"title":      title,
		"state":      state,
		"modifiedBy": modifiedBy,
		"modifiedOn": uint32(time.Now().Unix()),
	}
	return article.Update(dao.engine, values)
}

func (dao *Dao) DeleteArticle(id uint32) error {
	article := &model.Article{Model: &model.Model{ID: id}}
	return article.Delete(dao.engine)
}
