package routes

import (
	"resultanalyser/pkg/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

  authRouter := router.Group("/auth")
  authRouter.POST("/signup", controller.SignUp)
  authRouter.POST("/login", controller.Login)

  subjectRoute := router.Group("/subject")
  
  subjectRoute.GET("/", controller.ListSubjects)
  subjectRoute.GET("/code", controller.SubjectByCode)
  subjectRoute.POST("/create", controller.CreateSubject)
  subjectRoute.POST("/edit/:id", controller.UpdateSubject)
  subjectRoute.DELETE("/delete/:id", controller.DeleteSubjectController)

  tutorRoute := router.Group("/tutor")
  tutorRoute.GET("/", controller.ListAllTutors)
  tutorRoute.POST("/create", controller.CreateTutorController)

}

