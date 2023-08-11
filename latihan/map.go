package main

import (
	"fmt"
	"strings"
)



func main() {
	mobil := map[string]string{
		"honda":  "jazz",
		"toyota": "gr yaris",
	}

var key string = "HONDA"
var keyValue string = strings.ToLower(key) 

	fmt.Println(mobil)
	fmt.Println(mobil[keyValue])
	fmt.Println(mobil["toyota"])
	mobil["honda"]="crv" //change value for key honda
	mobil["nissan"]="grand livina" //append new key value to map
	fmt.Println(mobil) //print new map

}