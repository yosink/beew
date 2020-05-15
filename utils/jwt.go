package utils

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

func GenerateJwtToken() (string,error) {
	claim := Claims{
		Uid:            1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7*time.Hour).UnixNano(),
			Issuer:    "yookie",
		},
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return tokenObj.SignedString(beego.AppConfig.String("jwtkey"))
}

func ParseToken(token string) (*Claims,error)  {
	tokenObj, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return beego.AppConfig.String("jwtkey"), nil
	})
	if tokenObj != nil {
		if claims,ok := tokenObj.Claims.(*Claims);ok && tokenObj.Valid {
			return claims,nil
		}
	}
	return nil,err
}
