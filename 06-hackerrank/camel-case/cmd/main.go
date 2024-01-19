package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}

func camelcase(s string) int {
	result := 1
	for _, char := range s {
		if string(char) == strings.ToUpper(string(char)) {
			result++
		}
	}
	return result
}
