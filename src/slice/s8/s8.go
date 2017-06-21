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
