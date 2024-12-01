package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for input: ", input)
	fmt.Printf("Type of this rating %T", input)
}
