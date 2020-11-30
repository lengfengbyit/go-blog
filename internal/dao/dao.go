package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (dao *Dao) Begin() {
	dao.engine = dao.engine.Begin()
}

func (dao *Dao) Commit() error {
	dao.engine = dao.engine.Commit()
	return dao.engine.Error
}

func (dao *Dao) Rollback() {
	dao.engine = dao.engine.Rollback()
}

func (dao *Dao) Transaction(fn func(dao *Dao) error) error {
	dao.Begin()
	err := fn(dao)
	if err != nil {
		dao.Rollback()
		return err
	}
	return dao.Commit()
}
