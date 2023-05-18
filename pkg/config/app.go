package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var db *gorm.DB

func Connect(){
  d, err := gorm.Open(mysql.Open("spear:spear@/resultgo?charset=utf8&parseTime=True&loc=Local"))
  if err != nil {
    panic(err)
  }

  db = d
}

func GetDB() *gorm.DB {
  return db
}



