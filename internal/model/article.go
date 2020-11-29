package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
	Tags          []Tag  `gorm:"many2many:blog_article_tag"`
}

func (article Article) TableName() string {
	return "blog_article"
}

func (article *Article) First(db *gorm.DB) (*Article, error) {
	art := &Article{}
	err := db.Preload("Tags").First(art).Error
	return art, err
}

func (article *Article) Count(db *gorm.DB) (int, error) {
	var count int
	if article.Title != "" {
		db.Where("title like ?", "%"+article.Title+"%")
	}

	err := db.Model(article).Where("state = ?", article.State).
		Where("is_del = ?", 0).
		Count(&count).Error

	return count, err
}

func (article *Article) List(db *gorm.DB, pageOffSize, pageSize int) (articles []*Article, err error) {
	if pageOffSize > 0 && pageSize > 0 {
		db.Offset(pageOffSize).Limit(pageSize)
	}

	if article.Title != "" {
		db.Where("title like ?", "%"+article.Title+"%")
	}

	err = db.Preload("Tags").Model(article).Where("state = ?", article.State).
		Where("is_del = ?", 0).
		Find(&articles).Error

	return
}

func (article *Article) Create(db *gorm.DB) error {
	err := db.Attrs(article).FirstOrCreate(article, &Article{Title: article.Title}).Error
	return err
}

func (article *Article) Update(db *gorm.DB, values map[string]interface{}) error {
	return db.Model(&Article{}).
		Where("id = ? and is_del = ?", article.ID, 0).
		Update(values).Error
}

func (article *Article) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(article).Error
}
