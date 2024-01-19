package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n%v\n\n", caesarCipher("middle-Outz", 2))
}

func caesarCipher(s string, k int) string {
	var result string

	for _, ascii := range s {
		result += getCipheredChar(ascii, k)
	}

	return result
}

const (
	upperLettersBeginning = 65
	lowerLettersBeginning = 97
	alphabetLength        = 26
)

func getCipheredChar(ascii rune, k int) string {
	if isUpper(ascii) {
		shifted := (int(ascii)+k-upperLettersBeginning)%(alphabetLength) + upperLettersBeginning
		return string(rune(shifted))
	}

	if isLower(ascii) {
		shifted := (int(ascii)+k-lowerLettersBeginning)%(alphabetLength) + lowerLettersBeginning
		return string(rune(shifted))
	}
	return string(ascii)
}

func isUpper(ascii rune) bool {
	return (ascii >= upperLettersBeginning && ascii < upperLettersBeginning+alphabetLength)
}

func isLower(ascii rune) bool {
	return (ascii >= lowerLettersBeginning && ascii < lowerLettersBeginning+alphabetLength)
}
