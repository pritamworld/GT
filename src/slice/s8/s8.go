// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var n[] int

	// Appends numbers to the slice.
	for i := 1; i<=5; i++{
		n=append(n,i)
	}

	// Display each value in the slice.
	fmt.Println("\nInteger Slice Value")
	for i :=0; i<5; i++{
		fmt.Println(n[i])
	}

	// Declare a slice of strings and populate the slice with names.
	str := []string{"A", "B", "C", "D", "E", "F", "G"}

	// Display each index position and slice value.
	fmt.Println("\nString Slice Value")
	for i, s := range str{
		fmt.Println(i,s)
	}
	// Take a slice of index 1 and 2 of the slice of strings.
	str1 := str[1:3]
	
	// Display each index position and slice values for the new slice.
	fmt.Println("\nNew Slice Value")
	for i, s := range str1{
		fmt.Println(i,s)
	}
}
