package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func setName(t *Student, newName string) {
	t.name = newName
}
func main() {
	a := Student{"phj", 20, 99}
	fmt.Println("변경전", a)

	setName(&a, "newName")
	fmt.Println("변경후", a)
}
