package middleware

import (
	"fmt"
	"account/model"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authorize 授权
func (*Middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析token
		token, err := jwt.Parse(c.GetHeader("token"), func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return model.HmacSampleSecret, nil
		})

		if err != nil {
			log.Println(err)
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		log.Println(token)
		if ok {
			log.Println(claims["aud"])
		}
		c.Next()
	}
}
