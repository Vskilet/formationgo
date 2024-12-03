package maps

import (
	"warmup/users"
)

type Firm map[int]*users.User

func New() Firm{
	return make(Firm)
}

func (f Firm) Add(key int, user *users.User){
	f[key] = user
}

func (f Firm) Get(key int) *users.User{
	return f[key]
}

func (f Firm) DelUser(key int){
	delete(f, key)
}

func (f Firm) Delete(){
	f = nil
}