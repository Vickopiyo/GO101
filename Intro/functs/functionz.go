package main

import "fmt"
// type of parameter must be declared when passing it

func greeting(name string) {
	fmt.Printf("Hello  %s ,Functions!!", name)
}

// Types of parameters and return MUST always be declared...Only applicable to functions with return statement
// if more than two params of same type..declaration of type is done after the last params as seen below
// declaraion type of return is done before the body of function

func sum(a, b int) int {
	return a + b
}

// ELIPSIS   passed before type in the function as params means it can take any number of args of int type(variadic Function)
// otherwise it returns empty slice []
func sumVariadicFunc(numbers ...int) int {

	var total int
	for _, number := range numbers {
		total += number
	}
   
	fmt.Println(total)  
   // If rerurns goes above it return missing return.Return should always be the last line in a GO func 
	return total
}
   
// Returning more than one datatype like string and int therefore in return the two must be returned    
// ELIPSIS passed in a function always placed at the end of the params
 func varidaicFunc(number string, numeros ...int)  (string, int){   
    
   // OUR RETURNS VALUES...SPRINTF used when value(string) is passed to func 
   var sum int   
   message :=fmt.Sprintf("Hello %s , here is our Sum",number)       

   for _, num := range numeros { 
      sum +=num
   }         
   fmt.Println(sum)  

   return message, sum        
                            
 }

func main() {
	// How to call functions in GOLang
	// greeting("VICK")
	// r := sum( 43 , 45)
	// fmt.Println("Sum is :",r)
	sumVariadicFunc(3, 3, 3, 3, 3, 5)     

//   create two new variables even though names are similar but they are block scoped.
//   one variable takes the value of message and the other sum bcz fun returns two objects 

    info,sum :=varidaicFunc("Vick",5,6,7,7,8)    

    fmt.Println(info, sum)
                     
}
