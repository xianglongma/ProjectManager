package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xianglongma/ProjectManager/dao"
	"time"
)

var jwtKey = []byte("a_secret_secret")

type claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 发放token
func ReleaseToken(user *dao.User) (string, int64, error) {
	expirationTime := time.Now().Add(1 * 24 * time.Hour) // token有效时间是一天
	claims := &claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "maxianglong",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", 0, err
	}
	return tokenString, expirationTime.Unix(), nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *claims, error) {
	claims := &claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
