package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go")

	fmt.Printf("My name is %v \n", "Gautam Jha")

	CurrentTime := time.Now()

	fmt.Println("Current Date and Time ", CurrentTime.String())

}
