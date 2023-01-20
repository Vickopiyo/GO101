package main 
import "fmt"  

// GLOBAL VARIABLE    and  DEFER IN GO
// DEFER--pushes func call to a list..Used for clean up effects 
// DEFER executed after the func has returned 
// DEFER--executed in LIFO--lastInFirstOut 
var message string 
              
func fun1()  {   
	message= "Hello func 1 "
	fmt.Println(message)
	
}  
func func2()  {   
     message = "Hello func 2"
	 fmt.Println(message)
}

func main()  {       

 message= "Hello main func"
  
 fmt.Println(message)    
//    defering ,makes the func being called LAST in that bloc.     
 defer fun1() 
 func2()    

 fmt.Println(message)
	     
}    
     
             
