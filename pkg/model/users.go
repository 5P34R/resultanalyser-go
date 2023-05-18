package model

import (

	"gorm.io/gorm"
)


func CreateUser(user *User) *gorm.DB {
  res := db.Create(&user)
  return res
}

func FindUserByid(id float64) (*User, *gorm.DB){
  // fmt.Println("from models: ", id)
  var user User
  res := db.First(&user, "ID=?", id)
  return &user, res
}

func FindUser(email string) (*User, *gorm.DB) {
  var user User
  res := db.First(&user, "email=?", email)
  return &user, res 
}
