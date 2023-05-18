package model

import (
	"resultanalyser/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
  gorm.Model
  Email  string  `json:"email" gorm:"unique"`
  Password string `json:"password"`
}

type Subject struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required" gorm:"unique"`
	Semester string `json:"semester" binding:"required"`
}

type Tutor struct {
	gorm.Model
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID"`
	SubjectID uint    `json:"-"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Subject{}, &Tutor{}, &User{})
	// db.Create(&Subject{ Name: "Computer Networks", Code: "CST301", Semester: "S6" })
}

//  Tutor

