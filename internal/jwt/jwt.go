package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(web.AppConfig.DefaultString("jwt::key", "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"))

type Claims struct {
	Username string `json:"username"`
	UserId   int64  `json:"userid"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userName string, userId int64, userRole string) (string, error) {
	nowTime := time.Now()
	jwtTime := web.AppConfig.DefaultString("jwt::expirehours", "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o")
	t, err := strconv.Atoi(jwtTime)
	if err != nil {
		t = 24
	}
	expireTime := nowTime.Add(time.Duration(t) * time.Hour)

	claims := Claims{
		userName,
		userId,
		userRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    web.AppConfig.DefaultString("appname", "quince"),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {

		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			now := jwt.NumericDate{Time: time.Now()}
			if now.Unix() > claims.ExpiresAt.Unix() {
				return nil, errors.New("jwt.token.expired")
			} else {
				return claims, nil
			}
		}
	}
	return nil, err
}
