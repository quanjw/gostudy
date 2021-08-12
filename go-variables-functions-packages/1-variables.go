package main

import "fmt"

func main() {
	//声明变量并初始化变量
	/*
	var firstName string
	var firstName,lastName string
	var age int

	var (
	    firstName, lastName string
	    age int
	)

	var (
	    firstName string = "John"
	    lastName  string = "Doe"
	    age       int    = 32
	)

	var (
			firstName string = "John"
			lastName  string = "Doe"
			age       int    = 32
		)

	var (
	    firstName, lastName, age = "John", "Doe", 32
	)
	*/

	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Printf("%s  %s ’s age is %d",firstName,lastName,age)

	//声明常量
	const HTTPStatusOK = 200

	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

}
