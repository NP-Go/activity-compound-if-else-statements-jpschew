package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert your code here
	var year string
	fmt.Println("Enter a year: ")
	fmt.Scanln(&year)

	if year, _ := strconv.ParseInt(year, 10, 0); year%400 == 0 {
		fmt.Println("It is a leap year")
	} else if year%4 == 0 {
		fmt.Println("It is a leap year")
	} else {
		fmt.Println("It is not a leap year")
	}
}
