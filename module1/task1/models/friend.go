package models

import (
	"fmt"
	"log"
)

type Friend struct {
	Person
	Music []string
}

type IFriend interface {
	IPerson
	GetMusic() string
	AddMusic(music ...string) string
}

func NewFriend(name string, age uint8, educationalInstitution string, music ...string) IFriend {
	return &Friend{
		Person: *NewPerson(name, age, educationalInstitution),
		Music:  music,
	}
}

func (f Friend) String() string {
	return fmt.Sprintf(
		"Name: %s\nAge: %d\nEducation: %s\nMusic:  %s\n",
		f.Name, f.Age, f.EducationalInstitution, f.Music,
	)
}

func (f Friend) Describe() {
	log.Println(f)
}

func (f *Friend) AddMusic(music ...string) string {
	f.Music = append(f.Music, music...)
	return fmt.Sprintf("Updated music list: %v", f.Music)
}

func (f *Friend) GetMusic() string {
	return fmt.Sprintf("Music list: %v", f.Music)
}
