package dao

import (
	"gotour/blog-service/internal/model"
	"gotour/blog-service/pkg/app"
	"time"
)

func (dao *Dao)CheckTagExists(ids []string) bool {
	return model.Tag{}.CheckTagExists(dao.engine, ids)
}

func (dao *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(dao.engine)
}

func (dao *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(dao.engine, pageOffset, pageSize)
}

func (dao *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
			CreatedOn: uint32(time.Now().Unix()),
		},
	}
	return tag.Create(dao.engine)
}

func (dao *Dao) UpdateTag(id uint32, name string, state uint8, modifyBy string) error {
	tag := model.Tag{
		Model: &model.Model{ModifiedBy: modifyBy, ID: id},
	}

	values := map[string]interface{}{
		"state": state,
		"name": name,
		"modified_on": uint32(time.Now().Unix()),
	}
	return tag.Update(dao.engine, values)
}

func (dao *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(dao.engine)
}

