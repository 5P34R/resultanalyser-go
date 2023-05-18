package middleware

import (
	"fmt"
	"net/http"
	"resultanalyser/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ReqAuth(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
    // fmt.Println("git getting token")
		c.AbortWithStatus(http.StatusUnauthorized)
    return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("supersecret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    exp, ok := claims["exp"].(float64)
    if !ok {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
    }
		if time.Now().Unix() > int64(exp) {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
		}
    userId, ok := claims["sub"].(float64)
    if !ok {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
    }

    user, err := model.FindUserByid(userId)
    if err.Error != nil {
      c.AbortWithStatus(http.StatusUnauthorized)
      
      return
    }
    c.Set("user",  user)

		c.Next()
	} else {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error": err.Error(),
    })
    return
	}

}
