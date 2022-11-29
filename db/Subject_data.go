package db

import "github.com/ishtiaqhimel/go-api-server/model"

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
	Add(subject model.Subject)
	DeleteById(id string)
	UpdateById(id string, subject model.Subject)
}

func (r *SubjectRepo) Add(subject model.Subject) {
	r.Subjects = append(r.Subjects, subject)
}

func (r *SubjectRepo) GetAll() []model.Subject {
	return r.Subjects
}

func (r *SubjectRepo) DeleteById(id string) {
	for i, subject := range r.Subjects {
		if id == subject.Id {
			// Do we need the ordered value? - No
			r.Subjects[i] = r.Subjects[len(r.Subjects)-1]
			r.Subjects[len(r.Subjects)-1] = model.Subject{}
			r.Subjects = r.Subjects[:len(r.Subjects)-1]
			return
		}
	}
}

func (r *SubjectRepo) UpdateById(id string, sub model.Subject) {
	for i, subject := range r.Subjects {
		if id == subject.Id {
			r.Subjects[i] = sub
			return
		}
	}
}
