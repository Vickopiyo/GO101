package main

import (
	"fmt"   
)

func main() {

	// declaring variable using var
	// All imports and variables declared must be used  in GO programms
	// long way to declare variables
	var student1 string
	student1 = "Vick"
	var student string   

	student = "Opiyo"    
         
	// REASSIGNING VARIABLE         
	student1  = "VICTOR "        
	// Shortest way of declaring variable declaring and initialization
	// Variable declareed := are BLOCKED SCOPED----only accessed within the function 
	// integer--int while dec numbers are float64
	age := 26
	fmt.Println(student1, student, age)  

	// equivalent to console.log of JS
	var a int
	var b float64
	var c string
	var d bool

	//   OUTPUTS 0(INT), 0(FLOAT64), FALSE(string)
	fmt.Println(a, b, c, d)    

	// declaring constants in 
	const pi = 3.14 
	fmt.Println(pi)    
   
}
