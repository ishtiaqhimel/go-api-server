package db

import (
	"errors"
	"github.com/ishtiaqhimel/go-api-server/model"
)

// StudentRepo In memory Database for student
type StudentRepo struct {
	Students []model.Student
}

func NewStudent() *StudentRepo {
	return &StudentRepo{
		Students: []model.Student{
			{
				Id:        "1604099",
				FirstName: "Ishtiaq",
				LastName:  "Islam",
				Subjects: []model.Subject{
					{Id: "103", Title: "Data Structure", Code: "CSE-103"},
					{Id: "104", Title: "Algorithm", Code: "CSE-104"},
				},
			},
			{
				Id:        "1604098",
				FirstName: "ABC",
				LastName:  "DEF",
				Subjects: []model.Subject{
					{Id: "103", Title: "Data Structure", Code: "CSE-103"},
					{Id: "105", Title: "Computer Network", Code: "CSE-105"},
				},
			},
			{
				Id:        "1604097",
				FirstName: "XYZ",
				LastName:  "PQR",
				Subjects: []model.Subject{
					{Id: "104", Title: "Algorithm", Code: "CSE-104"},
					{Id: "106", Title: "Image Processing", Code: "CSE-106"},
				},
			},
		},
	}
}

type StudentService interface {
	GetAll() []model.Student
	Add(student model.Student) error
	DeleteById(id string) error
	UpdateById(id string, student model.Student) error
}

func (r *StudentRepo) Add(student model.Student) error {
	for _, stud := range r.Students {
		if stud.Id == student.Id {
			return errors.New("student with id " + student.Id + " already exists")
		}
	}
	r.Students = append(r.Students, student)
	return nil
}

func (r *StudentRepo) GetAll() []model.Student {
	return r.Students
}

func (r *StudentRepo) DeleteById(id string) error {
	for i, student := range r.Students {
		if id == student.Id {
			// Do we need the ordered value? - No
			r.Students[i] = r.Students[len(r.Students)-1]
			r.Students[len(r.Students)-1] = model.Student{}
			r.Students = r.Students[:len(r.Students)-1]
			return nil
		}
	}

	return errors.New("student with id " + id + " does not exist")
}

func (r *StudentRepo) UpdateById(id string, stud model.Student) error {
	for i, student := range r.Students {
		if id == student.Id {
			r.Students[i] = stud
			return nil
		}
	}

	return errors.New("student with id " + id + " does not exist")
}
