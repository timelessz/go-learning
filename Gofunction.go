package main

import (
	"fmt"
)

type sayer interface {
	walkway()
	changeAge()
}

type User struct {
	name    string
	age     int
	walkleg string
}

type Dog struct {
	name    string
	age     int
	walkleg string
}

func NewUser(name string, age int, walkleg string) User {
	return User{
		name:    name,
		age:     age,
		walkleg: walkleg,
	}
}

func NewDog(name string, age int, walkleg string) Dog {
	return Dog{
		name:    name,
		age:     age,
		walkleg: walkleg,
	}
}

func (u *User) walkway() {
	fmt.Printf("i am a person %#v ", u)
}

func (s *Dog) walkway() {
	fmt.Printf("i am a dog %#v ", s)
}

func (u *User) changeAge() {
	u.age = 24
}

func (u *Dog) changeAge() {
	u.age = 25
}

func BaseWalk(s sayer) {
	s.walkway()
}

func main() {
	u := NewUser("赵兴壮", 1, "两条腿")
	d := NewDog("jsonen", 4, "四条腿")
	BaseWalk(&u)
	BaseWalk(&d)
	u.changeAge()
	d.changeAge()
	BaseWalk(&u)
	BaseWalk(&d)
}
