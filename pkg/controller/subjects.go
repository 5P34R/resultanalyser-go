package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"resultanalyser/pkg/model"
)

func ListSubjects(c *gin.Context) {
	c.JSON(200, gin.H{
		"Subjects": model.GetAllSubject(),
	})
}

func CreateSubject(c *gin.Context) {
	var subject model.Subject
	err := c.BindJSON(&subject)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{

			"error":   "VALIDATIONERROR",
			"message": err.Error(),
		})
		return
	}

	res, err := model.CreateSubject(&subject)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, &res)
}

func SubjectByCode(c *gin.Context) {

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "VALIDATIONERROR",
			"message": "code is not given",
		})
		return
	}

	subject,_, err := model.GetSubjectByCode(code)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, subject)
}

func UpdateSubject(c *gin.Context) {
	code := c.Param("id")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "code is not provided",
		})
	}
	var newSubject model.Subject
	err := c.BindJSON(&newSubject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	old, db , err := model.GetSubjectByCode(newSubject.Code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	if newSubject.Name != "" {
		old.Name = newSubject.Name

	}

	if newSubject.Code != "" {
		old.Code = newSubject.Code
	}
	if newSubject.Semester != "" {

		old.Semester = newSubject.Semester
	}

  db.Save(&old)
  
	c.JSON(http.StatusOK, &old)
}


func DeleteSubjectController(c *gin.Context){
  code := c.Param("id")
  model.DeleteSubject(code)
  c.JSON(http.StatusOK, gin.H{
    "message":"deleted successfully",
  })
}
