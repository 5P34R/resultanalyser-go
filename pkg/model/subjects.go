package model

import (
	"errors"

	"gorm.io/gorm"
)


func GetAllSubject() []Subject {
	var subjects []Subject
	db.Find(&subjects)
	return subjects
}

func CreateSubject(subject *Subject) (*Subject, error) {
	res,_, err := GetSubjectByCode(subject.Code)
	if err == nil {
		return res, errors.New("Already Exists")
	}
	db.Create(&subject)
	return subject, nil
}

func GetSubjectByCode(code string) (*Subject, *gorm.DB, error) {
	var subject Subject
	res := db.Where("CODE=?", code).Find(&subject)
	if subject == (Subject{}) {
		return &subject,res, errors.New("Doesnt Exist")
	}
	return &subject,res, nil
}

func DeleteSubject(code string) Subject {
  var subject Subject
  db.Where("CODE=?", code).Delete(&subject)
  return subject
}


