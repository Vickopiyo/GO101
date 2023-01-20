package main

import (
	"fmt"
	"strings"
)

// CLOSURES
//Func takes in an int and returns a string
func repeat(n int)  func (name string) string {  


	// This anonymous func closes over the func repeat cz it obtains a value(n) from it to accomplished or executed 
	// it makes the func stay alive because it utilizes the n-integer     
	return func(name string)  string{    
		return strings.Repeat(name, n) 
	}	
}   

func main()  {  
//   first instance it takes in a int but it returns a func which takes in str thus requires to be called again
   repeat3 := repeat(3)  
   //   Prints only memory address cz it returns a func which is a variable(memory allocated) 
   fmt.Println(repeat3("vick"))
   
}         