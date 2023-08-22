package main

import (
	"catfacts/catfacts"
	"log"
)

// Main function that we call from the command line
func main() {
	s := catfacts.NewClient("https://catfact.ninja/", "Mozilla/5.0")

	fact, err := s.ListCatFact()
	CheckError(err)
	log.Print(fact.String())

	facts, err := s.GetNumberOfCatFacts(103)
	CheckError(err)
	log.Print((*facts)[0].String())

}

// Helper function that errors out if we get an error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
