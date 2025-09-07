package parcticev2

import (
	"fmt"
	"math"
)

type MakeAppear func(string) string

func main() {

	innerFn := makeAppear("Hello world")

	fmt.Println(innerFn("MOBIN"))

	fmt.Println(factorial(5))

	fmt.Println(fibonacci(6))

	fmt.Println(math.Round(avg(10, 20)))
	fmt.Println(math.Round(avg(10, 20, 54, 79, 30)))
	fmt.Println(max(10, 20, 54, 79, 30))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...)) // Use slice splitting
}

func makeAppear(base string) MakeAppear {
	return func(innerStr string) string {
		return innerStr + " " + base
	}
}

func factorial(n int) int {
	// factNumber := 1
	// for i := 1; i <= n; i++ {
	// 	factNumber *= i
	// }

	// return factNumber

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func fibonacci(n int) int {

	prev, curr := 0, 1

	for i := 2; i < n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}

	return curr

}

func avg(nums ...float64) float64 {

	totalMoney := 0.0
	totalNumbers := len(nums)

	for i := 0; i < totalNumbers; i++ {
		totalMoney += nums[i]
	}
	return totalMoney / float64(totalNumbers)
}

func max(values ...int) int {

	maxValue := values[0]

	for _, value := range values {

		if maxValue < value {
			maxValue = value
		}
	}

	return maxValue
}

func sum(nums ...int) int {
	sumValue := 0
	for _, value := range nums {
		sumValue += value
	}

	return sumValue
}

// 5 * 24 = 120
// 4 * (3 * 2) = 6
// 3 * 2
// 2 * 1
// 1 * 1
// ret -> 1
