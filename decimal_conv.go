package main

import "fmt"

func main() {

	var decimalNumber int

	fmt.Printf("Enter a decimal number: ")
	fmt.Scanln(&decimalNumber)

	fmt.Printf("Binary Number = %b\n", decimalNumber)
	fmt.Printf("Octal Number = %o\n", decimalNumber)
	fmt.Printf("Hexadecimal Number = %x\n", decimalNumber)

}
