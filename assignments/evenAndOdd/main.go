package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		typeNumber := "even"
		if number%2 == 0 {
			typeNumber = "odd"
		}
		fmt.Println(number, "is", typeNumber)
	}
}
