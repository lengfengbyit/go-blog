package dao

import "gotour/blog-service/internal/model"

func (dao *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(dao.engine)
}
