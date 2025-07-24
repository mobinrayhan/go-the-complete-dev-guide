package main

import (
	"fmt"
	"reflect"
)

type Mover interface {
	Move()
}

type Speaker interface {
	Speak()
}

type Robot interface {
	Mover
	Speaker
}

type Worker interface {
	Work()
}

type Employ struct {
	Name string
}

func (e Employ) Work() {
	fmt.Println(e.Name, "is working.")
}

func (e Employ) Extra() {
	fmt.Println(e.Name, "is doing extra tasks!")
}

func main() {

	// a := mobile.Android{
	// 	Name: "Samsung S24 Ultra",
	// }

	// performRobotTasks(a)
	// describeTypes("Hello World")
	// describeTypes(45645)
	// describeTypes(false)

	// typeOfValue(42)
	// fmt.Println()

	// typeOfValue("Hello")
	// fmt.Println()

	// typeOfValue([]string{"a", "b", "c"})
	// fmt.Println()

	// typeOfValue(map[string]int{"a": 1, "b": 2})

	w := Employ{
		Name: "mobin",
	}

	printWorkData(w)
}

func printWorkData(work Worker) {
	work.Work()
	// work.Extra()
}

func performRobotTasks(r Robot) {
	r.Move()
	r.Speak()
}

// func describe(value interface{}) {

// 	switch v := value.(type) {
// 	case int:
// 		fmt.Println("Print INT")
// 	case string:
// 		fmt.Println("Print STRING")
// 	case float64:
// 		fmt.Println("Print STRING")
// 	default:
// 		fmt.Println(v)
// 	}
// }

func describeTypes(value interface{}) {

	switch v := value.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	case float64:
		fmt.Println(v)
	case float32:
		fmt.Println(v)
	case bool:
		fmt.Println(v)
	default:
		fmt.Println(v)
	}
}

func typeOfValue(value interface{}) {

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	fmt.Println("Type:", t.Name()) // e.g., "int", "string"
	fmt.Println("Kind:", t.Kind()) // e.g., "int", "slice", "struct"
	fmt.Println("Value:", v)
}
