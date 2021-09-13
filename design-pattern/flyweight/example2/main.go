package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(FullName string) *User {
	return &User{FullName: FullName}
}

var allNames []string

type User2 struct {
	names []uint8
}

func NewUser2(FullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(FullName, " ")

	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}

	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}
func main() {
	john := NewUser("John D")
	jinny := NewUser("Jinny D")
	alsoJinny := NewUser("Jinny C")
	fmt.Println("Memory taken by users:",
		len([]byte(john.FullName))+
			len([]byte(jinny.FullName))+
			len([]byte(alsoJinny.FullName)))

	john2 := NewUser2("John D")
	jinny2 := NewUser2("Jinny D")
	alsoJinny2 := NewUser2("Jinny C")
	fmt.Println(john2.names)
	fmt.Println(jinny2.names)

	// fmt.Println(john2.FullName())
	// fmt.Println(jinny2.FullName())
	// fmt.Println(alsoJinny2.FullName())
	fmt.Println("Memory taken by users:",
		len([]byte(john2.names))+
			len([]byte(jinny2.names))+
			len([]byte(alsoJinny2.names)))

	fmt.Println(allNames)
}
