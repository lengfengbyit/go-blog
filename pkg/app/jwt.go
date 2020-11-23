package app

import (
	"github.com/dgrijalva/jwt-go"
	"gotour/blog-service/global"
	"gotour/blog-service/pkg/helper"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"App_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	expireTime := time.Now().Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    helper.EncodeMd5(appKey),
		AppSecret: helper.EncodeMd5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 解析 token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)

		// 验证token时间，是否合法
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
