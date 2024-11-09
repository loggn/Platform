package pkg

import (
	"golang.org/x/crypto/bcrypt"
)
/*
	功能：密码加密
	描述：加密用户密码
	创建时间：2024-11-6
*/
func HashPassword(password string) (string, error) {
	// bcrypt 默认会使用 10 的加密强度，可以根据需要调整
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

/*
	功能：密码对比
	描述：比较用户密码
	创建时间：2024-11-6
*/
func CheckPasswordHash(password, hash string) bool {
	// 比较用户输入的密码和存储的哈希值
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}