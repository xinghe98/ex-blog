package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// 自定义需要加密的参数
type Myclaims struct {
	Username string `json:"username"`
	UserKey  string `json:"userkey"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}

var Secret = []byte(viper.GetString("secret.key"))

//生成token
func GenToken(username string, userkey string, id uint) (string, error) {
	claims := Myclaims{
		username,
		userkey,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //设置token过期时间
			Issuer:    "xinghe",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

//解析token
func ParseToken(tokenstr string) (*jwt.Token, *Myclaims, error) {
	claims := &Myclaims{}
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return token, claims, err
}
