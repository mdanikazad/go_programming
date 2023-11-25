package main

import "fmt"

func main() {
    
	var year int
	fmt.Println("Enter a year: ")
	fmt.Scanln(&year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(year, " is a leap year")
	} else {
		fmt.Println(year, " is not a leap year")
	}

}
