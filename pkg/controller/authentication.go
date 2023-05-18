package controller

import (
	"net/http"
	"resultanalyser/pkg/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	type requestBody struct {
		Email    string `json:"email"`
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
	c.JSON(http.StatusAccepted, gin.H{
		"user": user,
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

  c.JSON(http.StatusOK, gin.H{
    "message": user,
  })
  
}
