package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (auth Auth) TableName() string {
	return "blog_auth"
}

func (auth Auth) Get(db *gorm.DB) (Auth, error) {
	db = db.Where(
		"app_key = ? and app_secret = ? and is_del = ?",
		auth.AppKey,
		auth.AppSecret,
		0)

	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
