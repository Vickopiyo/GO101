package main  
import "fmt"

func main()  {   
//   Go has Not,OR  and And logical operators 
// Syntax   NOT --  ! 
    // NOT 
	// fmt.Println(!true)   
//   AND -- && ---returns true if both sides of equal is true 
   fmt.Println( true &&  false)
//    OR --returns true when either side of eqaution returns true 

  fmt.Println(false || false)  

  b := 2 
  r := b==2 && !(b > 4)
//   returns  false bcz second part negates true to false 
  fmt.Println(r)           
  
// IF statements   
// Program to check if you are an adult 
var age int      
var adult bool
fmt.Print("How old are you?")  

fmt.Scanln(&age)   

if age >= 18   {      
	if age <= 35{
		adult = true    
	   fmt.Printf("its %t you are  still a  youth!", adult)   
	}else if age >= 36 && age <= 50  {
		fmt.Println("You are nologer a youth!")	
	}else if  age >= 80  && age < 90 {
		  fmt.Println("You are an octogerian!!")     
	} else{
		fmt.Println("Hello great grand pa!!!")  
	}   
	      
}else{   
	adult = false  
	fmt.Printf("Age must be an integer!")     
}          
}