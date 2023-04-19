package main

import (
	"fmt"
)

func main() {
	s1 := []int{5, 3, 1, 9, 2, 7}
	s2 := []int{2, 4, 5, 7, 8, 9, 15, 20, 25, 30, 32, 35}
	fmt.Println(linearSearch(s1, 9))
	fmt.Println(binarySearch(s2, 30))
}
