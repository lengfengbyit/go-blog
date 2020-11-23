package model

import "github.com/jinzhu/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (tag Tag) TableName() string {
	return "blog_tag"
}

func (tag Tag)CheckTagExists(db *gorm.DB, ids []string) bool {
	var count int
	db.Model(&tag).Where("id in (?)", ids).Count(&count)
	return len(ids) == count
}

func (tag Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if tag.Name != "" {
		db = db.Where("name = ?", tag.Name)
	}

	db = db.Where("state = ?", tag.State)
	err := db.Model(&tag).
		Where("is_del=?", 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (tag Tag) List(db *gorm.DB, pageOffset, pageSize int) (tags []*Tag, err error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if tag.Name != "" {
		db = db.Where("name=?", tag.Name)
	}
	db = db.Where("state = ?", tag.State)
	err = db.Where("is_del=?", 0).
		Find(&tags).Error

	return
}

func (tag Tag) Create(db *gorm.DB) error {
	return db.Create(&tag).Error
}

func (tag Tag) Update(db *gorm.DB, values map[string]interface{}) error {
	err := db.Model(&Tag{}).
		Where("id = ? AND is_del=?",tag.ID, 0).Updates(values).Error
	return err
}

func (tag Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?",
		tag.ID, 0).Delete(&tag).Error
}
