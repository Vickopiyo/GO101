package main  

import "fmt"   


func main()  {     
// LOOPS AND ITERATION   
// for only with no conditions loops infinitely  
// INFINITE LOOP
// for  {
// 	fmt.Printf("Loop at:")
// }   

//    num := 1000   
//    times := 0 
//    for num > 0 { 
// 	  num /=10 
// 	  times++
//    }   
//    fmt.Println("Number of times ",times)   


// COMMON SYNTAX FOR  FOR  LOOP   
// A program writing 1 to 10 
 for i := 1; i <= 10; i++ {    
	if i % 2 ==0  {
		fmt.Printf("we are at number %d snd it's an even number ", i)
	}
	  fmt.Printf("We are at number %d   though its an odd number \n",i)
 }   
   
//  A program for decrementing    
 
    for i := 10; i >= 0 ; i-- {   
		 fmt.Printf("We are on a down roll at %d \n ",i)   
	}  


     


   



}