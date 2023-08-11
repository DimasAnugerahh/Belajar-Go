package main

import "fmt"

//implementasi penggunaan fungsi sebagai parameter pada fungsi lainnya


type Filter func(string)string 
//implementasi typeDeclaration sebagai alias pada fungsi untuk mempersingkat parameter pada fungsi


func helloName(name string,filter Filter)  {
	filteredName:=filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "kotor" {
		return "..."
	} else {
		return name
	}
}

func main() {
helloName("kotor",spamFilter)
}