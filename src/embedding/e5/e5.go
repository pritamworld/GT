// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the match method for hockey, call into the match method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import ("strings"
		"fmt")


// matcher defines the behavior required for performing matching.
type matcher interface {
	match(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	teamName string
	cityName string
}

// match checks the value for the specified term.
func (s sport) match(searchTerm string) bool {
	return strings.Contains(s.teamName, searchTerm) || strings.Contains(s.cityName, searchTerm)
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.
type hockey struct {
	sport
	country string
}

// match checks the value for the specified term.
func ( h hockey ) match(searchTerm string) bool {

	// Make sure you call into match method for the embedded sport type.

	// Implement the search for the new fields.
	return h.sport.match(searchTerm) || strings.Contains(h.country, searchTerm)
}

func main() {

	// Define the term to match.
	term := "Toronto"

	// Create a slice of matcher values and assign values
	// of the concrete hockey type.
	matchers := []matcher{
		hockey{sport{"Cricket", "Gujarat"}, "INDIA"},
		hockey{sport{"Hockey", "Toronto"}, "CANADA"},
		sport{"Hockey", "Toronto"},
	}

	// Display what we are trying to match.
	fmt.Println("Try to match for:", term)

	// Range of each matcher value and check the term.
	for _, m := range matchers {
		if m.match(term) {
			fmt.Printf("\nFOUND: %+v", m)
		}
	}
}
