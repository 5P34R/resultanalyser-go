package controller

import (
	"fmt"
	"net/http"
	"resultanalyser/pkg/model"

	"github.com/gin-gonic/gin"
)

type TutorList struct {
  Name string
  Phone string
}


func ListAllTutors(c *gin.Context){
  var tutors []model.Tutor
  user, _ := c.Get("user")
  fmt.Println(user)
    tutors = model.ListTutors()

  var response []gin.H 

  for _, tutor := range tutors{
    response = append(response, gin.H{
      "name": tutor.Name,
      "phone":tutor.Phone,
    })
  }

  if len(response) <= 0 {
    c.JSON(http.StatusOK, gin.H{
      "message": "Tutors list is empty",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "tutors": response,
  })
}

func CreateTutorController(c *gin.Context) {
   
   type TutorReq struct {
     Name string `json:"name"`
     Phone string `json:"phone"`
     SubjectId string `json:"subjectcode"`
   }

  type User struct {
    Email string
    Password string
  }

  // user, _ := c.Get("user")
  
  var tutorReq TutorReq
  err := c.Bind(&tutorReq)
  if err != nil {
    c.JSON(http.StatusNotAcceptable, gin.H{
      "error": "VALIDATIONERROR",
      "message":err.Error(),
    })
    return
  }

   subject, _, err := model.GetSubjectByCode(tutorReq.SubjectId)
   if err != nil{
     c.JSON(http.StatusNotAcceptable, gin.H{
       "error": "Subject doesn't exists",
     })
     return
   }
  

   tutor := model.Tutor{
     Name: tutorReq.Name,
     Phone: tutorReq.Phone,
     Subject: *subject,
   }
   
   model.CreateTutors(&tutor)

  c.JSON(http.StatusOK, gin.H{
    "tutor": tutor,
  })

}


