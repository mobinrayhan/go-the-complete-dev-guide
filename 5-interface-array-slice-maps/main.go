package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Product struct {
	Id    string
	Title string
	Price float64
}

func main() {

	hobbiesArr := [3]string{"Codding", "Listing Songs", "Eating A Lot Chicken"}

	fmt.Println(hobbiesArr)
	fmt.Println(hobbiesArr[0])

	secondAndThird := hobbiesArr[1:3]
	fmt.Println(secondAndThird)

	firstAndSecondEle := hobbiesArr[0:2]
	firstAndSecondEleV2 := hobbiesArr[:2]

	fmt.Println(firstAndSecondEle, firstAndSecondEleV2)
	trickySlice := firstAndSecondEleV2[0:3]

	fmt.Println(trickySlice)
	// firstAndSecondEle

	courseGoals := []string{"Practice Codding", "Being Great Dev"}

	fmt.Println(courseGoals)
	courseGoals[1] = "Being Soft Dev Like A PRO"
	courseGoals = append(courseGoals, "Third Goal So its slice like pro")
	fmt.Println(courseGoals)

	id, err := GenerateRandomId()

	if err != nil {
		fmt.Println("Error generating ID:", err)
		return
	}

	listProducts := []Product{{Title: "Mobile 1", Id: id, Price: 500}, {Title: "Mobile 1sdfasdfasdf", Id: id, Price: 5200}}

	fmt.Println(listProducts)
	listProducts = append(listProducts, Product{Title: "Mobile ss1", Id: id, Price: 52200})
	fmt.Println(listProducts)
}

func GenerateRandomId() (string, error) {
	bytes := make([]byte, 16)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list

// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)

// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.

// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
