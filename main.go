package main

import "fmt"

// Constant Expressions and Iota

const (
	one = iota * 2 // will be 0*2
	two            //will be 1*2
)

const (
	//Iota will be reset to 0 in this code constant segment.
	a = iota
	b
)

func main() {

	//Iota and Constant Expressions
	//  param two is not defined so go uses constant expression for the second variable.
	fmt.Println(one, two)

	// Iota resets the count if and when defined in a different constant block
	fmt.Println(a, b)

	// Declaring variables(Value Types)
	var i int
	i = 4
	fmt.Println(i)
	fmt.Println("hooray")

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Mark"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 5)
	fmt.Println(c)

	right, left := real(c), imag(c)
	fmt.Println(right, left)

	//Pointers Type. ( Holds the address of the location in memory that holds the information)

	// Initialize the pointer
	var secondName *string = new(string)
	//De-reference the pointer using the asterisk
	*secondName = "james"
	fmt.Println(secondName)
	//Dereference the pointer to access actual values
	fmt.Println(*secondName)

	//Address of Operator to change the value of variable.
	carName := "Toyota"
	fmt.Println(carName)

	//Use Address of operator (&) to get pointer
	pointer := &carName
	fmt.Println(pointer)

	carName = "Benz"
	fmt.Println(carName)
	fmt.Println(pointer)

	// Constants
	const pi = 3 // Implicit constant
	// Type checked at runtime. so mathematical expressions like the ones below will work perfectly
	fmt.Println(3 + 4.55)
	fmt.Println(3 + 4)
	const age int = 23

	//fmt.Println(age + 3.444) // wont work because type was already specified
	fmt.Println(float32(age) + 3.444095498598549595699865989065090969809560050680065) // will work because type has been manually converted

}
