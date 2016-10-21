// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

// Add imports.
import "fmt"

// Declare the speaker interface with a single method called speak.
type speaker interface{
	speak()
}

// Declare an empty struct type named english.
type english struct{}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"
func (english) speak(){
	fmt.Println("Hello world");
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named speak for the chinese type
// using a pointer receiver. "你好世界"
func (c *chinese) speak(){
	fmt.Println("你好世界")
}

//Created By: Pritesh Patel
func (c chinese) hello(){
	fmt.Println("Chinese : 你好世界")
}

//Created By: Pritesh Patel
func (c *chinese) noHello(){
	fmt.Println("No : 你好世界")
}

// sayHello accepts values of the speaker type.
func sayHello( s speaker ) {
	s.speak()
	// Call the speak method from the speaker parameter.
}

func main() {

	// Declare a variable of the interface speaker type
	// set to its zero value.
	var sp speaker

	// Declare a variable of type english.
	var e english

	// Assign the english value to the speaker variable.
	sp = e

	// Call the speak method against the speaker variable.
	sp.speak()

	// Declare a variable of type chinese.
	var c chinese

	// Assign the chinese pointer to the speaker variable.
	sp = &c

	// Call the speak method against the speaker variable.
	sp.speak()
	
	//Created By: Pritesh Patel
	c.hello()
	c.noHello()

	// Call the sayHello function with new values and pointers
	// of english and chinese.
	sayHello(english{})
	//sayHello(c)
	//ERRRRRORRR
	//interface/e6/e6.go:75: cannot use c (type chinese) as type speaker in argument to sayHello:
	//chinese does not implement speaker (speak method has pointer receiver)
   
	sayHello(&english{})
	sayHello(&chinese{})
}
