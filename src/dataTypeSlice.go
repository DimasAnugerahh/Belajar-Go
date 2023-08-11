package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	// var output []int = slice[3:5]

	fmt.Println(slice[3:5])

	slice = append(slice,12,1333)
	fmt.Print(slice)
}