package model

type Specialty struct {
	id          int
	description string
}

func NewSpecialty(description string) Specialty {
	return Specialty{
		description: description,
	}
}

func (s Specialty) GetId() int {
	return s.id
}

func (s *Specialty) SetId(id int) {
	s.id = id
}

func (s Specialty) GetDescription() string {
	return s.description
}

func (s Specialty) SetDescription(description string) {
	s.description = description
}
