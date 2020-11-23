package service

import (
	"gotour/blog-service/internal/model"
	"gotour/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}


func (ser *Service) CountTag(param *CountTagRequest) (int, error) {
	return ser.dao.CountTag(param.Name, param.State)
}

func (ser *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return ser.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (ser *Service) CreateTag(param *CreateTagRequest) error {
	return ser.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (ser *Service) UpdateTag(param *UpdateTagRequest) error {
	return ser.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (ser *Service) DeleteTag(param *DeleteTagRequest) error {
	return ser.dao.DeleteTag(param.ID)
}
