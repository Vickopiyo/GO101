package main

import (

	"fmt"
	"strconv"
	"strings"
)

func sum(expression string) int {

	var results int
	//COnverts expression into array.Separator is the operator(+)
	values := strings.Split(expression, "+")

	fmt.Println(values)
	// Loop through the created array
	for _, value := range values {
		// Initialize a new  var which  matches the elements in the array and convert them into int
		// NIL---equivalent to undefined(JS)--unintialized
		//  NIL--means a zero value for pointers, interfaces,maps and slices and channels.Only supported in this datatypes
		// Equivalent of nil for int is 0(default), float(0.0), string(""),

		num, error := strconv.Atoi(value) 
		// checks if there is any error in the values in the slice 
		if error != nil {  
			// Handles any error      
			fmt.Println(error)   
			// populates the first argument if it is correct nad throws if there are more errors 
			fmt.Println("Please enter a number!")       
			fmt.Println("Please check if you included + ")   
		}else{
			results += num        
		}   

	}

	// FOR PERFOMANCE!!       

	// for i := 0; i < len(values); i++ {
	// 	num, _ :=strconv.Atoi(values[i])
	//      results += num
	// }
	// fmt.Println(results)

	return results
}

func main() {

	var expression string

	var result int
	fmt.Print("=>")
	//  Here you passed values and operation to be peformded(+)
	fmt.Scanln(&expression)
	result = sum(expression)
	fmt.Printf("Result is %d \n", result)

}
