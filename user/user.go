package user

import (
	"math/rand"
	"time"
)

type User struct {
	Id int
	Name string
	Firstname string
	Birthday time.Time 
}

func Create(name string, firstname string, birthday time.Time) *User {
	return &User{
		Id: rand.Intn(100),
		Name: name,
		Firstname: firstname,
		Birthday: birthday,
		//CreatedAt: time.Now(),
	}
}
