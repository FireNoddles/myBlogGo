package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

var (
	TokenExpired     error = errors.New("Token已过期,请重新登录")
	TokenNotValidYet error = errors.New("Token无效,请重新登录")
	TokenMalformed   error = errors.New("Token不正确,请重新登录")
	TokenInvalid     error = errors.New("这不是一个token,请重新登录")
)

var JwtKey = []byte("kjs1dh41a")

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		JwtKey,
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (j *JWT) SetToken(username string, pwd string) (string, error) {
	setClaim := MyClaims{
		username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaim)
	token, err := reqClaim.SignedString(j.JwtKey)
	if err != nil {
		log.Error("SetToken err", err)
		return "", err
	}

	return token, err
}

func (j *JWT) CheckToken(token string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		log.Error("jwt.ParseWithClaims err [%v]", err)
		return nil, err
	}

	if t != nil {
		if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
			return claims, nil

		}

		return nil, TokenInvalid
	}

	return nil, TokenInvalid

}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := SUCCSE

		if tokenHeader == "" {
			code = ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": GetErrMsg(code),
			})
			c.Abort()
			return
		}

		CheckToken := strings.Split(tokenHeader, " ")
		if len(CheckToken) == 0 {
			code = ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(CheckToken) != 2 || CheckToken[0] != "Bearer" {
			code = ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := &JWT{
			[]byte(JwtKey),
		}
		claim, err := j.CheckToken(CheckToken[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  ERROR,
				"message": err,
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claim)
		c.Next()

	}
}
