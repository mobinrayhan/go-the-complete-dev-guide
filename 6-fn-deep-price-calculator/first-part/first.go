package firstpart

import "fmt"

type OperateCB func(int, int) int
type AddGenerator func(int) int

func main() {
	greet := func(name string) {
		fmt.Printf("Hello %s, Good Morning!\n", name)
	}
	greet("Mobin")
	greet("Himon")
	greet("Robin")

	multiply := operate(10, 20, func(a, b int) int {
		return a * b
	})

	addResult := operate(20, 30, add)

	fmt.Println(addResult)
	fmt.Println(multiply)

	retFn := makeAdder(10)
	fmt.Println(retFn(20))

	next := counter()

	fmt.Println(next())
	fmt.Println(next())

	fmt.Println(next())

	fmt.Println(next())

	fmt.Println(next())

	fmt.Println(next())

	fmt.Println(next())

}

func operate(a, b int, cb OperateCB) int {
	return cb(a, b)
}

func add(a, b int) int {
	return a + b
}

func makeAdder(x int) AddGenerator {
	return func(givenNumber int) int {
		return x + givenNumber
	}
}

func counter() func() int {
	count := 0
	return func() int {
		// count = count + 1
		count++
		return count
	}
}
