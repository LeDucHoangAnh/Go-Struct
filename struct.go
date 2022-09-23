package main

import "fmt"

//khoi tao
type Student struct {
	id   int
	name string
	info StudentInfo
}
type StudentInfo struct {
	class string
	email string
	age   int
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
	// student2 := Student{456, "Mary"}
	// fmt.Println(student2)
	// fmt.Println(student2.id)
	// fmt.Println(student2.name)

	// var student3 Student = struct {
	// 	id   int
	// 	name string
	// }{
	// 	777,
	// 	"Ken",
	// }

	// fmt.Println(student3)

	// anonymous struct : struct vo danh
	var anonymous = struct {
		email string
		age   int
	}{
		"kaiz@gmail.com",
		17,
	}

	fmt.Println(anonymous)

	//pointer tro toi 1 struct
	// pointerStruct := &Student{
	// 	000,
	// 	"Leo",
	// }
	// fmt.Println(&pointerStruct)
	// fmt.Println(pointerStruct.id)
	// fmt.Println(pointerStruct.name)
	// anonymous field
	type noName struct {
		string
		int
	}

	n := noName{
		"Layla",
		18,
	}
	fmt.Println(n)

	//Struct long struct - Nested Struct
	student_A := Student{
		id:   123,
		name: "many",
		info: StudentInfo{
			class: "001",
			email: "@gmail.com",
			age:   28,
		},
	}
	fmt.Println(student_A)

	//compare 2 struct
	type struct1 struct {
		id   int
		name string
		// info map[int]int
	}

	s1 := struct1{
		1,
		"A",
		// map[int]int{
		// 	0: 1,
		// },
	}

	s2 := struct1{
		1,
		"A",
		// map[int]int{
		// 	0: 1,
		// },
	}

	if s1 == s2 {
		fmt.Println("s1 = s2")
	} else {
		fmt.Println("s1 != s2")
	}
}
