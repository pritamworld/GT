// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type ball struct{
	name string
	atBats int
	hits int
}

// Declare a method that calculates the batting average for a player.
func (b *ball) average() float64 {
	avg := float64(b.hits) / float64(b.atBats)
	//avg := float64(b.hits / b.atBats)
	return avg
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	balls := []ball{
		{"A",10,100},
		{"B", 20,50},
	}

	// Display the batting average for each player in the slice.
	for _, s := range balls{
		fmt.Println(s.name,s.average())
	}
}
