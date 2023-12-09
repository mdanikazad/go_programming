//basics of variable

package main

import "fmt"

func main() {

	// var declaration
	//fmt.Println("Hello World")
	var fullName, country string
	var age int
	var gpa float32

	// var initialization
	fullName = "Onik Azad"
	country = "Bangladesh"
	age = 32
	gpa = 5

	fmt.Println(fullName, "is a student")
	fmt.Println(fullName, "is ", age, "years old")
	fmt.Println(fullName, "has got ", gpa, "/5 in SSC")
	fmt.Println(fullName, "originally from ", country)

}
