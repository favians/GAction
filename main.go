package main

import (
	"fmt"
)

func main() {
	val := Hello("Main")

	fmt.Println(val)

	myStruct := MyStruct{
		Name:     "Vian",
		Password: "12345678",
	}
	structToMap := StructTomap(myStruct)
	fmt.Println(structToMap)
	fmt.Println(structToMap)

}
