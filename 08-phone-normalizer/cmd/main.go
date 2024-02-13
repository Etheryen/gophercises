package main

import (
	"08-phone-normalizer/database"
	"08-phone-normalizer/env"
	"08-phone-normalizer/utils"
	"fmt"
	"log"
	"time"
)

var phoneNumbers = [...]string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
}

func main() {
	start := time.Now()

	fmt.Println("Loading .env...")
	handleErr(env.Load())

	fmt.Println("Initializing db...")
	handleErr(database.Init())

	fmt.Println("Populating db...")
	handleErr(database.Populate(phoneNumbers[:]))

	fmt.Println("Reading all rows...")
	rows, err := database.ReadAll()
	handleErr(err)
	utils.PrintArray(rows)

	fmt.Println("Normalizing phone numbers...")
	handleErr(database.NormalizeAll())

	fmt.Println("Reading final rows...")
	rows, err = database.ReadAll()
	handleErr(err)
	utils.PrintArray(rows)

	fmt.Println("Took:", time.Since(start).Round(time.Millisecond))
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
