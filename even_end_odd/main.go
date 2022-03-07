package main

import "fmt"

func main() {

	// fmt.Print("Enter number : ")
	// var n int
	// fmt.Scanln(&n)
	// /*  Conditional Statement if .... else ........     */
	// if n%2 == 0 {
	// 	fmt.Println(n, "is Even number")
	// } else {
	// 	fmt.Println(n, "is Odd number")
	// }

	numbers := []int{}
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "is Even number")
		} else {
			fmt.Println(num, "is Odd number")
		}
	}
	// fmt.Println(numbers)
}
