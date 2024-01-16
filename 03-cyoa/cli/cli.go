package cli

import (
	"03-cyoa/types"
	"fmt"
	"strconv"
)

func Start(storyMap map[string]types.StoryArc, firstArcName string) {
	currentArc := storyMap[firstArcName]
	for len(currentArc.Options) > 0 {
		printArcDetails(currentArc)
		chosenOption := getNumber("Chosen option", len(currentArc.Options))
		nextArc := currentArc.Options[chosenOption-1].Arc
		currentArc = storyMap[nextArc]
	}
	printArcDetails(currentArc)
	fmt.Println("The End")
}

func printArcDetails(arc types.StoryArc) {
	fmt.Printf("\n--- %v ---\n\n", arc.Title)
	for _, paragraph := range arc.Paragraphs {
		fmt.Println(paragraph)
	}
	fmt.Println()
	for i, option := range arc.Options {
		fmt.Printf("%v. %v\n", i+1, option.Text)
	}
	if len(arc.Options) > 0 {
		fmt.Println()
	}
}

func getNumber(message string, maxNum int) int {
	fmt.Printf("%v: ", message)
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Something went wrong when scanning number, try again...")
		return getNumber(message, maxNum)
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Incorrect integer, try again...")
		return getNumber(message, maxNum)
	}

	if num < 1 || num > maxNum {
		fmt.Println("Incorrect option, try again...")
		return getNumber(message, maxNum)
	}

	return num
}
