package main

import (
	"fmt"

	piscine "piscine/cameltosnakecase"
)

func main() {
	fmt.Println(piscine.CamelToSnakeCase("HelloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("helloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("camelCase"))
	fmt.Println(piscine.CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(piscine.CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(piscine.CamelToSnakeCase("hey2"))
}
