package main

import (
	"errors"
	"fmt"
)        
// Divder func returns two things...result or error
func division(num1, num2 int)  (int, error){
	
if num2 == 0 {
	// errors pkg is imported with func New to output a guide on error 
   return 0, errors.New("Divider is 0!")    
}else{   
	// here calculation is done and error is nil(undefined--uninitialized)   
	return num1/num2, nil
}
}

func main()  {
	// ERRORHANDLING    
   //  we assign two variables to the two output of division()
	result,err :=division(8,5)   
	// check condition of output  
	// NO errors      
	if err == nil { 
		fmt.Println("Result is :",result)
	}else{
		fmt.Println(err)
	}                 
         
}   