package db

import (
	"errors"
	"github.com/ishtiaqhimel/go-api-server/model"
)

// SubjectRepo In memory database for subjects
type SubjectRepo struct {
	Subjects []model.Subject
}

func NewSubject() *SubjectRepo {
	return &SubjectRepo{
		Subjects: []model.Subject{
			{Id: "103", Title: "Data Structure", Code: "CSE-103"},
			{Id: "104", Title: "Algorithm", Code: "CSE-104"},
			{Id: "105", Title: "Computer Network", Code: "CSE-105"},
			{Id: "106", Title: "Image Processing", Code: "CSE-106"},
		},
	}
}

type SubjectService interface {
	GetAll() []model.Subject
	Add(subject model.Subject) error
	DeleteById(id string) error
	UpdateById(id string, subject model.Subject) error
}

func (r *SubjectRepo) Add(subject model.Subject) error {

	for _, sub := range r.Subjects {
		if subject.Id == sub.Id {
			return errors.New("subject with id " + subject.Id + " already exists")
		}
	}
	r.Subjects = append(r.Subjects, subject)
	return nil
}

func (r *SubjectRepo) GetAll() []model.Subject {
	return r.Subjects
}

func (r *SubjectRepo) DeleteById(id string) error {
	for i, subject := range r.Subjects {
		if id == subject.Id {
			// Do we need the ordered value? - No
			r.Subjects[i] = r.Subjects[len(r.Subjects)-1]
			r.Subjects[len(r.Subjects)-1] = model.Subject{}
			r.Subjects = r.Subjects[:len(r.Subjects)-1]
			return nil
		}
	}
	return errors.New("subject with id " + id + " does not exist")
}

func (r *SubjectRepo) UpdateById(id string, sub model.Subject) error {
	for i, subject := range r.Subjects {
		if id == subject.Id {
			r.Subjects[i] = sub
			return nil
		}
	}
	return errors.New("subject with id " + id + " does not exist")
}
