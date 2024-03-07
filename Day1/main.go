package main

import "fmt"



func main() {

	fmt.Println("hello world")

	//variable decleration

	// var name string = "webflux"
	// var nameTwo = "favour"

	// var matric int = 20205714

	// var matric1 int8 =  123

	// namegwo := "luigi"

	// can't use the shorthand outside of a function.

	//int, float, double

	// var score float32 = 2.323

	//for float the size has to be specified.

	//float32 =
	//float64 =


	// fmt.Println(nameTwo, name, matric, matric1, namegwo, score)

	//Arrayyyyyyyyyyyyys and Slices

	var Ages [5]int = [5]int {24, 23, 25, 30, 27}

	for i := 0; i < len(Ages); i++ {
		fmt.Println(Ages[i])
	}

	// Slices uses Arrayyyyyyyyyyyyys under the hood

	


}
