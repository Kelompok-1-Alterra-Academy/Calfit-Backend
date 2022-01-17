package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTMyClaims struct {
	Id               int
	Email            string
	Member           bool
	OperationalAdmin bool
	Superadmin       bool
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTMyClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(Id int, Email string, Member bool, OperationalAdmin bool, Superadmin bool) (string, error) {
	claims := JWTMyClaims{
		1,
		Email,
		Member,
		OperationalAdmin,
		Superadmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(jwtConf.ExpiresDuration)).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}

func GetUser(c echo.Context) *JWTMyClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTMyClaims)
	return claims
}
