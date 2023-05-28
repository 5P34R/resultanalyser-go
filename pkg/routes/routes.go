package routes

import (
	"resultanalyser/pkg/controller"
	"resultanalyser/pkg/middleware"

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
  tutorRoute.GET("/", middleware.ReqAuth, controller.ListAllTutors)
  tutorRoute.POST("/create", middleware.ReqAuth, controller.CreateTutorController)

}

