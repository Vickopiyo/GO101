package main  

import "fmt"    

func main()  {     

	// CLOSURES AND ANONYMOUS FUNCTIONS...Anonymous func aka lambda func aka func literal..Closure and anonymous are quite similar though closure is an instance of a func.    
	// Anonymous function--function without name..Called immediately at the end of it 
	// Anonymous functions are used  for short term execution and passed into HOF--higher order functions
	//   
	// SHOULDNOT BE WRITTEN AS STAND ALONE AS SEEN BELOW 
    //   func ()  {
	// 	fmt.Println("I am an anonymous function!!")
	//   }()     
    //  ANONYMOUS func MUST  BE  ASSIGNED TO A VARIABLE 
	myFunc := func (name string) string {
		return fmt.Sprintf("I %s am an anonymous function!!", name)
	  }       

	//   ALSO Anonymous func can be INVOKED IMMEDIATELY and variable passed somewhere else 
	// myFunc := func (name string) string {
	// 	return fmt.Sprintf("I %s am an anonymous function!!", name)
	//   }("Denno")       
    // passing func without calling returns memory address   
	  fmt.Println(myFunc("Vick"))        

	  
}   