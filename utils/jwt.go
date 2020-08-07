package utils

import (
	"aries/config/setting"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 载荷
type CustomClaims struct {
	Username string `json:"username"`
	UserImg  string `json:"user_img"`
	jwt.StandardClaims
}

// 常量
var (
	TokenExpired     = errors.New("令牌已过期")
	TokenNotValidYet = errors.New("令牌未激活")
	TokenMalformed   = errors.New("令牌格式有误")
	TokenInvalid     = errors.New("无效的令牌")
	SignKey          = "aries-open-source-blog" // 签名
)

//创建一个 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(GetSignKey()),
	}
}

// 获取 SignKey
func GetSignKey() string {
	return SignKey
}

// 设置 SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// 创建 Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新 Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
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
			Duration(setting.Config.Server.TokenExpireTime)).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
