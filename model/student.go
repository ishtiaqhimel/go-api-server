package model

type Student struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Subjects  []Subject `json:"subjects"`
}
