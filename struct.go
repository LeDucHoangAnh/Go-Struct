package main

import "fmt"

//khoi tao
type Student struct {
	id   int
	name string
}

func main() {
	// named
	student1 := Student{
		id:   123,
		name: "Kai",
	}

	fmt.Println(student1)
	fmt.Println(student1.id)
	fmt.Println(student1.name)
	//
	student2 := Student{456, "Mary"}
	fmt.Println(student2)
	fmt.Println(student2.id)
	fmt.Println(student2.name)

	var student3 Student = struct {
		id   int
		name string
	}{
		777,
		"Ken",
	}

	fmt.Println(student3)

	// anonymous struct : struct vo danh
	var anonymous = struct {
		email string
		age   int
	}{
		"kaiz@gmail.com",
		17,
	}

	fmt.Println(anonymous)

}
