package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// type declaration
type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]

}

func main() {
	users := []User{
		{"Rangga", 21},
		{"Dwi", 35},
		{"Mahendra", 27},
		{"Hendra", 25},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
