package main

import "fmt"

func main() {
	test := []int{1, 2, 3}

	test = append(test, 5)

	fmt.Println(test)
}
