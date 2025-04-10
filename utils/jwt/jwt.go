package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secret_crect")

type Claims struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

// 生成 JWT token
func GenerateToken(username string, role int) (string, error) {
	// 设置过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		username,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// ParseToken 解析 JWT token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}
	return nil, err
}
