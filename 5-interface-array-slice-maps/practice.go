package main

import (
	"fmt"

	"mobin.dev/slice/person"
)

type Age int

func main() {
	favInts := []int{10, 20, 30}

	for index, value := range favInts {
		switch index {
		case 0:
			fmt.Println("First : ", value)
		case 1:
			fmt.Println("Second : ", value)
		case 2:
			fmt.Println("Third : ", value)
		default:
			fmt.Println("NO match Found!")
		}
	}

	// groceryList := map[string]int{"apples": 5, "bananas": 3, "milk": 1}

	// fmt.Println(groceryList, (groceryList["bananas"]))

	groceryList := map[string]int{"apples": 2, "oranges": 3}
	fmt.Println(groceryList)

	groceryList["bananas"] = 3

	fmt.Println(groceryList)

	groceryList["apples"] = 3
	fmt.Println(groceryList)

	delete(groceryList, "oranges")
	fmt.Println(groceryList)

	// 🟡 4. Maps vs Structs

	peopleMap := []map[string]string{
		{"Name": "Mobin"},
		{"Name": "HImon"},
	}

	peopleStruct := []person.Person{
		{Name: "Alice", Age: 25, Country: "USA"},
		{Name: "Bob", Age: 30, Country: "Canada"},
	}

	fmt.Println(peopleMap)
	fmt.Println(peopleStruct)

	// 🟡 5. Using The Special make Function

	capSlice := make([]int, 0, 5)
	capMapFloat := make(map[string]int)

	fmt.Println(capSlice)
	fmt.Println(capMapFloat)

	capSlice = append(capSlice, 10, 20, 5, 45, 02, 23, 23, 23)
	fmt.Println(capSlice)

	capMapFloat["apple"] = 54
	capMapFloat["mango"] = 544

	fmt.Println(capMapFloat)

	makingMap := make(map[string]string)
	makingMap["userName"] = "Mobin"
	makingMap["email"] = "movin@test.com"

	makingMap["role"] = "admin"

	// for index, value := range makingMap {
	// 	fmt.Print(index, value)
	// }

	for key, value := range makingMap {
		fmt.Println(key, value)
	}
	// keys := []string{"userName", "email", "role"}

	// for i := 0; i < len(keys); i++ {
	// 	fmt.Println(makingMap[keys[i]])
	// }

	userData := map[string]string{"userName": "Mobin", "email": "mobin.dev"}
	userData["role"] = "Admin"
	fmt.Println(userData)

	johnAge := map[string]Age{"Joffhn": 222}

	printInfo(johnAge)

	aSliceOfFiveNumbers := []int{10, 20, 30, 40, 50}
	studentsGrade := map[string]int{"Alice": 85, "Bob": 91}

	for i := 0; i < len(aSliceOfFiveNumbers); i++ {
		fmt.Println(aSliceOfFiveNumbers[i])
	}

	for key, value := range studentsGrade {
		fmt.Printf("%s Scored %d\n", key, value)
	}
}

func printInfo(ages map[string]Age) {
	for key, value := range ages {
		fmt.Printf("%s is %d years old", key, value)
	}
}
