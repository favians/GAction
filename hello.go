package main

import (
	"fmt"

	"github.com/fatih/structs"
)

type MyStruct struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Hello(value string) string {
	return fmt.Sprintf("Hello From: %s", value)
}

func StructTomap(value MyStruct) map[string]interface{} {
	return structs.Map(value)
}
