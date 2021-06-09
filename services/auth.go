package services

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thedoor-dev/back/configs"
)

type _claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

var salt string = "qunmjbricimfvs"

func genJWTToken(username string) (string, error) {
	claims := _claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().
				Add(60 * 60 * 24 * 3 * time.Second).Unix(),
			Issuer: "the door",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(salt))
}

func parseToken(token string) (*_claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&_claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(salt), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*_claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Token")
		cliaim, err := parseToken(token)
		if err != nil {
			ResponseError(c, http.StatusUnauthorized, codeNoRight)
			c.Abort()
			log.Println(token)
			return
		}
		// cliaim.ID
		if cliaim.Username == configs.Conf.AdminName {
			c.Set("uname", cliaim.Username)
			c.Next()
		} else {
			c.Abort()
			return
		}
	}
}

func Signin(c *gin.Context) {
	data := struct {
		AdminName   string `json:"name"`
		AdminPasswd string `json:"pawd"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	if data.AdminName != configs.Conf.AdminName || data.AdminPasswd != configs.Conf.AdminPasswd {
		ResponseError(c, http.StatusProxyAuthRequired, codeRefuse)
		return
	}
	token, err := genJWTToken(data.AdminName)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"name":  data.AdminName,
		"img":   configs.Conf.Img,
		"token": token,
	})
}
