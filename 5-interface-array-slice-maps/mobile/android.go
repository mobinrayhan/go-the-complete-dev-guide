package mobile

import "fmt"

type Android struct {
	Name string
}

// Android implements Move() method
func (a Android) Move() {
	fmt.Println(a.Name, "is moving forward.")
}

// Android implements Speak() method
func (a Android) Speak() {
	fmt.Println(a.Name, "says: Hello, human!")
}
