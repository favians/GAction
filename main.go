package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	val := Hello("Main")

	fmt.Println(val)

	myStruct := MyStruct{
		Name:     "Vian",
		Password: "12345678",
	}
	structToMap := StructTomap(myStruct)
	mapB, _ := json.Marshal(structToMap)
	fmt.Println(string(mapB))
	Hello := "Hello From ACTIONS Here Is Your Data: " + string(mapB)
	fmt.Fprintf(w, Hello)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	handleRequests()
}
