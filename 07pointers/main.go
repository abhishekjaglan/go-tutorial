package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	var ptr *int
	fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	ptr = &myNumber
	fmt.Println("Value of pointer is: ", ptr)
	fmt.Println("Value of pointer pointing to is: ", *ptr)

	var secondPtr = &myNumber
	fmt.Println("Value of pointer is: ", secondPtr)
	fmt.Println("Value of pointer pointing to is: ", *secondPtr)

	var thirdPtr *int = &myNumber
	fmt.Println("Value of pointer is: ", thirdPtr)
	fmt.Println("Value of pointer pointing to is: ", *thirdPtr)

	*thirdPtr = *thirdPtr * *thirdPtr

	fmt.Println("Value of first pointer: ", *ptr)
	fmt.Println("Value of second pointer: ", *secondPtr)
	fmt.Println("Value of third pointer: ", *thirdPtr)
	fmt.Println("Value of myNumber: ", myNumber)
}
