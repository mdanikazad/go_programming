package main

import "fmt"

func main() {

	// var declaration using constant
	const (
		fullName = "Onik Azad"
		country  = "Bangladesh"
		age      = 32
		gpa      = 5
	)

	// printing variable
	fmt.Printf("%v is a student\n", fullName)
	fmt.Printf("%v is %v years old\n", fullName, age)
	fmt.Printf("%v has got %v/5 in SSC\n", fullName, gpa)
	fmt.Printf("%v is originally from %v", fullName, country)

}
