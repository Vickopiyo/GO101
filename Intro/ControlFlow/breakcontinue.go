package main   
import "fmt"          


func main()  {       

	// BREAK--- key word used for breaking  loop 
	// CONTINUE--code in the if block runs when condition is met and loops proceeds 
  for i := 0; i <= 10; i++ {  
    // if loops reeaches at 5 it executes code block and allow loop to continue running
	if i==5  {
		fmt.Println("Don't be scared.We are not terminating loop!")  
		continue
	}
    //  condition to break our loop
	    if  i==7  {
			// code to be executed when loop breaks 
			// code will run normally if break key word is not placed 
           fmt.Println("Sorry we had to break the loop!") 
           break
		}  

  fmt.Println(i)   

  }
}