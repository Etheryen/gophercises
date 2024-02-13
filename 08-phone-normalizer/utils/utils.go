package utils

import "fmt"

func StrToAnySlice(strSlice []string) []any {
	var anySlice []any

	for _, s := range strSlice {
		anySlice = append(anySlice, s)
	}

	return anySlice
}

func PrintArray[T any](arr []T) {
	fmt.Println("[")
	for _, el := range arr {
		fmt.Printf("  %v\n", el)
	}
	fmt.Println("]")
}
