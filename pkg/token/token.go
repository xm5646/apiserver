/**
 * 功能描述: json web token
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	// Header中不存在 `Authorization`
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero. ")
)

type ContextToken struct {
	ID       uint64
	Username string
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*ContextToken, error) {
	ctx := &ContextToken{}

	//解析token
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// 从请求中解析token
func ParseRequest(c *gin.Context) (*ContextToken, error) {
	header := c.Request.Header.Get("Authorization")

	// 从配置文件加载jwt secret
	secret := viper.GetString("server.jwt_secret")

	if len(header) == 0 {
		return &ContextToken{}, ErrMissingHeader
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

func Sign(ctx *gin.Context, c ContextToken, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("server.jwt_secret")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err = token.SignedString([]byte(secret))
	return
}
