package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secre")

// CustomClaims 自定义声明结构体，用于添加额外的声明
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(username string) (string, error) {
	// 设置过期时间为一小时
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建声明
	claims := &CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 创建并签名 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken 验证 JWT 令牌
func VerifyToken(tokenString string) (*CustomClaims, error) {
	// 解析 JWT
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证并返回声明
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
