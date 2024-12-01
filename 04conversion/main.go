package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Rating of our pizza from you is:", input)
	fmt.Printf("Type of input is: %T\n", input)

	floatRating, err := strconv.ParseFloat(strings.TrimSpace(input), 32) // used TrimSpace becasue string input is- "4\n", this removes the \n
	//intRating, err := strconv.ParseInt(input, 0, 32)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", floatRating+1)
	}
}
