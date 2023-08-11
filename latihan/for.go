package main

import "fmt"


func main() {
	counter := 0
	for i := counter; i < 10; i++ {
		fmt.Println(i)
	}

	//for range

	slice :=[...]int{1,2,3,4,5};
	mobil:=map[string]string{
		"toyota":"alphard",
		"honda":"s2000",
		"nissan":"gtr",
	}

	for key, value := range mobil {//map
		fmt.Println("key ",key,"value ",value)
	}
	for key, value := range slice {//slice
		fmt.Println("key ",key,"value ",value)
	}

//kalau index tidak d perlukan pada perulangan, bisa d ganti underscore
	fmt.Println()
	for _, value := range mobil {
		fmt.Println(value)
	}
}