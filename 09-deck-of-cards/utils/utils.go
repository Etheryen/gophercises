package utils

import "fmt"

func PrintArray[T any](arr []T, cols int) {
	fmt.Println("[")
	for i := 0; i < len(arr); {
		fmt.Print(" ")
		for j := 0; j < cols; j++ {
			if i >= len(arr) {
				break
			}
			fmt.Printf(" %v,", arr[i])
			i++
		}
		fmt.Println()
	}
	fmt.Println("]")
}
