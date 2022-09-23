package main

import "fmt"

//khoi tao
type Student struct {
	id   int
	name string
}

func main() {
	student1 := Student{
		id:   123,
		name: "Kai",
	}

	fmt.Println(student1)
}
