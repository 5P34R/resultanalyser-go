package model

import (

	"gorm.io/gorm"
)


func CreateUser(user *User) *gorm.DB {
  res := db.Create(&user)
  return res
}

func FindUser(email string) (*User, *gorm.DB) {
  var user User
  res := db.First(&user, "email=?", email)
  return &user, res 
}
