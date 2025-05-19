package container

import (
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
}

func (u People) Generate() {
	fmt.Println("People...")
}
