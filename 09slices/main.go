package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go")

	var fruitList = []string{"Apple", "Peach", "Mango"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Lichi")
	fmt.Println(fruitList)
	fmt.Println(fruitList[0], fruitList[4])

	fruitList2 := append(fruitList[1:3])

	fmt.Println(fruitList2)

	fruitList3 := append(fruitList[:3])
	fmt.Println(fruitList3)

	highScores := make([]int, 4) // declaring slice without initializing it (kind of like dynamic memory allocation)

	highScores[0] = 134
	highScores[1] = 934
	highScores[2] = 334
	highScores[3] = 434

	highScores = append(highScores, 555, 666, 243) // alocates additional memory of 3 to already existing of 4

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println("Just for the sake of maintaining streak")
}
