package basics

import "fmt"

// package level
var c, phyton, java bool

func Variables() {
	// function level
	var i int
	fmt.Println("=====================variables===================")
	fmt.Println(i, c, phyton, java)
	
	// variables with initializers
	var j, k, name = 2, 4, "luis"
	// := short assigment statement
	age := 6
	fmt.Println("===========variables with initializers===========")
	fmt.Println(j, k, name, age)
}