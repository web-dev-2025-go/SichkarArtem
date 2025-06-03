package model

import (
	"fmt"
)

type Me struct {
	Person
	Hobbies string
}

type IMe interface {
	IPerson
	NotChangeAge(age uint8) uint8
	ChangeAge(age uint8) uint8
	GetHobbies() string
}

func NewMe(name string, age uint8, educationalInstitution string, hobbies string) IMe {
	return &Me{
		Person: Person{
			name,
			age,
			educationalInstitution,
		},
		Hobbies: hobbies,
	}
}

func (m Me) GetHobbies() string {
	return m.Hobbies
}

func (m Me) NotChangeAge(age uint8) uint8 {
	m.Age = age
	// not change age in struct
	return m.Age
}

func (m *Me) ChangeAge(age uint8) uint8 {
	m.Age = age
	// change age
	return m.Age
}

func (m Me) String() string {
	return fmt.Sprintf(
		"Name: %s\nAge: %d\nEducation: %s\nHobbies:  %s\n",
		m.Name, m.Age, m.EducationalInstitution, m.Hobbies,
	)
}

func (m Me) Describe() {
	fmt.Println(m)
}
