package pkg

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

/*
	功能：生成token
	描述：生成token，成功后返回 Token，有效期24小时
	创建时间：2024-11-6
*/
var jwtSecret = []byte("yourSecretKey")

func GenerateToken(user User) (string, error) {
	// 设置过期时间为 24 小时
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"sub":  user.ID,
		"exp":  expirationTime.Unix(),
	}
	// 使用 jwt.SigningMethodHS256 签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成 token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}