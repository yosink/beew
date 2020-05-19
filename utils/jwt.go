package utils

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

type AuthToken struct {
	Token    string `json:"token"`
	ExpireAt string `json:"expire_at"`
}

func GenerateJwtToken(uid int) (*AuthToken, error) {
	expired := time.Now().Add(7 * time.Hour).Unix()
	claim := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expired,
			Issuer:    "yookie",
		},
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := beego.AppConfig.String("jwtkey")
	token, err := tokenObj.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}
	return &AuthToken{
		Token:    token,
		ExpireAt: time.Unix(expired, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	if tokenObj != nil {
		if claims, ok := tokenObj.Claims.(*Claims); ok && tokenObj.Valid {
			return claims, nil
		}
	}
	return nil, err
}
