package main

import "fmt"

func main() {

	// var initialization using type inference
	fullName := "Onik Azad"
	country := "Bangladesh"
	age := 32
	gpa := 5

	fmt.Println(fullName, "is a student")
	fmt.Println(fullName, "is ", age, "years old")
	fmt.Println(fullName, "has got ", gpa, "/5 in SSC")
	fmt.Println(fullName, "originally from ", country)

}
