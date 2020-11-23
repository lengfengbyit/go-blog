package service

import "github.com/pkg/errors"

type AuthRequest struct {
	AppKey    string `form:"appKey" binding:"required"`
	AppSecret string `form:"appSecret" binding:"required"`
}

func (ser *Service) CheckAuth(param *AuthRequest) error {
	auth, err := ser.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("Auth info does not exits.")
}
