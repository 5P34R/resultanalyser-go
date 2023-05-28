package controller

import (
	"net/http"
	"resultanalyser/pkg/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

  type ResponseBody struct {
    ID uint `json:"id"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"CreatedAt"`
    UpdatedAt time.Time `json:"UpdatedAt"`
  }

	var reqBody requestBody
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


  hash, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Failed to generate password",
    })
    return
  }

  user := model.User{Email: reqBody.Email, Password: string(hash)}
  res := model.CreateUser(&user)
  if res.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Failed to create user",
    })
    return 
  }
  respones := ResponseBody{
    ID: user.ID,
    Email: user.Email,
    CreatedAt: user.CreatedAt,
    UpdatedAt: user.UpdatedAt,
  }
	c.JSON(http.StatusAccepted, gin.H{
		"user": respones,
	})
}

func Login(c *gin.Context) {
  type requestBody struct {
    Email string `json:"email"`
    Password string `json:"password"`
  }

  var reqBody requestBody

  err := c.BindJSON(&reqBody)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  user, res := model.FindUser(reqBody.Email)

  if res.Error != nil {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error": "Invalid email or password",
    })
    return
  }

  erro := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

  if erro != nil {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error": "Invalid email or password",
    })
    return
  }

  // Create a new token object, specifying signing method and the claims
// you would like it to contain.
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	  "sub": user.ID,
	  "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
  })

  // Sign and get the complete encoded token as a string using the secret
  tokenString, err := token.SignedString([]byte("supersecret"))

  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error":err.Error(),
    })
    return
  }

  c.SetSameSite(http.SameSiteLaxMode)
  c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

  c.JSON(http.StatusOK, gin.H{
    "token": tokenString,
  })
  
}

