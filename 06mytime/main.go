package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("15:04:05 Monday 02-01-2006"))

	createdDate := time.Date(2023, time.April, 26, 12, 0, 0, 0, time.UTC)
	fmt.Println("Created Data: ", createdDate)
	fmt.Println(createdDate.Format("15:04:05 Monday 02-01-2006"))
}
