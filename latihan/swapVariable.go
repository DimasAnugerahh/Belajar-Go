package main

import "fmt"

func main() {
	a := 23
	b := 3

	a, b = b, a

	fmt.Println(a,b)
}