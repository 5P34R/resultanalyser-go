
package model


func ListTutors() []Tutor {
  var tutors []Tutor
  db.Find(&tutors)
  return tutors
}

func CreateTutors(tutor *Tutor)*Tutor {
  db.Create(&tutor)
  return tutor
}
