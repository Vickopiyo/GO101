package main  
import "fmt"

func main()  {   
//   Go has Not,OR  and And logical operators 
// Syntax   NOT --  ! 
//    NOT 
	// fmt.Println(!true)   
//   AND -- && ---returns true if both sides of equal is true 
   fmt.Println( true &&  false)
//    OR --returns true when either side of eqaution returns true 

  fmt.Println(false || false)  

  b := 2 
  r := b==2 && !(b > 4)
  
//   returns  false bcz second part negates true to false 
  fmt.Println(r)


	
}