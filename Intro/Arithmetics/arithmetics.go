package main   
import "fmt" 

func main()  {    

	// REASSIGNING OF VARIABLES DOES NOT REQUIRE the : colon
	     
   a := 20 
   b := 30  

   result := a + b  
   fmt.Println("Result is: ",result)
      
   result = b - a 
   fmt.Println("Results Two is: ", result)    

   result = a * b 
   fmt.Println("Result Three is:", result)

}