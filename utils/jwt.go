package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// CustomClaims 载荷
type CustomClaims struct {
	Username string `json:"username"`
	UserImg  string `json:"user_img"`
	jwt.StandardClaims
}

// 常量
var (
	ErrTokenExpired     = errors.New("令牌已过期")
	ErrTokenNotValidYet = errors.New("令牌未激活")
	ErrTokenMalformed   = errors.New("令牌格式有误")
	ErrTokenInvalid     = errors.New("无效的令牌")
	SignKey             = "aries-open-source-blog" // 签名
)

// NewJWT 创建一个 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(GetSignKey()),
	}
}

// GetSignKey 获取 SignKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 设置 SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 创建 Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token
func (j *JWT) RefreshToken(tokenString string, tokenExpireTime int) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SigningKey, nil
		})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Second * time.
			Duration(tokenExpireTime)).Unix()
		return j.CreateToken(*claims)
	}

	return "", ErrTokenInvalid
}
