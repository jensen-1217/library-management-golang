package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 为给定的纯文本密码生成哈希密码
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
