package main

import "fmt"

func main() {

	// *--deferencing operator used to declare pointer variable and access the value stored in the address
	// &--  address operator used to return the address of a variable or to access address of a variable to a pointer

	var x int = 578      

	// var y = 567      

	var p *int

	p = &x        

	fmt.Println("Value stored in x =", x )  
	fmt.Println("Address of x  =", &x )          
	fmt.Println("Value stored in p =", p )      
                   
}        