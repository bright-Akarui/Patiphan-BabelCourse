package container

import (
	"fmt"
)

type User struct {
	Name   string
	Course []string
}

func (u User) Generate() {
	fmt.Println("Generating for user", u.Name)
}
