package main

import (
	"fmt"
)

// asyncronus, defer akan di jalankan setelah seluruh statement d kerjakan

func secondFunction() {
	fmt.Println("WORLD")
}

func firstFunction(){
	defer secondFunction()
	fmt.Println("Hello")
}

func main() {
firstFunction()
}