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

   result = b / a 
    // Even though accurate rssults was 1.5..it prints out  1 beacuse result was initialized as int  not float64
   fmt.Println("Results four is :", result )       
   
//    TO get results i  float64 always  declare var to be float64 and depednded varibale to be in float64 format too 

var div float64 = 6.0 / 4.0   
    fmt.Println("Division ",div)   

	// Modulus--only applicable to to integers 
	div =  6 % 4     
	fmt.Println("Modulus :",div)
       
}        