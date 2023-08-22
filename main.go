package main

import (
	"catfacts/catfacts"
	"log"
)

// Main function that we call from the command line
func main() {
	c := catfacts.NewClient(
		"https://catfact.ninja/",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
	)

	fact, err := c.ListCatFact()
	CheckError(err)
	log.Print(fact.String())

	facts, err := catfacts.GetNumberOfCatFacts(c, 103)
	CheckError(err)
	log.Print((*facts)[0].String())

}

// Helper function that errors out if we get an error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
