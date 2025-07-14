package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	user := Person{
		FirstName: "Mobin",
		LastName:  "Rayhan",
		Age:       25,
	}

	fmt.Print(user)
	birthday(&user)
	// birthday(&user)
	// birthday(&user)
	// birthday(&user)
	// birthday(&user)
	// birthday(&user)

	fmt.Println(user)
	// user.FirstName = "Mobin"
	// user.LastName = "Rayhan"
	// user.Age = 22
}

func birthday(p *Person) {
	p.Age += 1
}
