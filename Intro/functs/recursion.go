package main  
         
import "fmt"  

  func factorial(n int)  int{  
    //  TAIL RECURSION--a func calt itself calculates value and sends it down the hierachy
	// KEY BENEFIT--- always faster than normal recursion   
    //   BASE condition...if met recursion stops  
	   if n == 0 {
		return 1
	   }  
    //  outputs 5,4,3,2,1 ...stops when it reaches 0        
	   fmt.Println(n)                        
	//first instance it gets 5 then it call then get 4 then 3 then 2 then 1  ()  
    // After base condition is met then last output assumes value of n and     
	// outputs 1,2,6,24,120  
	 //  tail recursive call     
    f := n * factorial(n-1)                                   
    fmt.Println(f)                                   
    return f                                     
	           
  }

func main()  {   
    
  factorial(5)    
	
}                 