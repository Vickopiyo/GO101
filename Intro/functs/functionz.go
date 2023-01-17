package main  
import "fmt"   

// type of parameter must be declared when passing it  

func greeting(name string)  {
	fmt.Printf("Hello  %s ,Functions!!",  name)      
}              
            
// Types of parameters and return MUST always be declared...Only applicable to functions with return statement    
// if more than two params of same type..declaration of type is done after the last params as seen below 
// declaraion type of return is done before the body of function     

func sum (a, b int) int {  

       return a + b 
}       
 
func main()  {   
	// How to call functions in GOLang
   greeting("VICK")	  
   r := sum( 43 , 45)   
   fmt.Println("Sum is :",r)                
}          





   