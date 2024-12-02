package main

import "fmt"

func main() {
	fmt.Println("Welcome to array tutorial in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list at 2 index is: ", fruitList[2])

	fmt.Println("Length of fruitlist is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", len(vegList))
}
