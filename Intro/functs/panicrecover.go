package main

import (
	"fmt"  
	"os"
)


func main()  {       
	
	// Throws this if file is not found  
 defer func ()  {  
	if error :=recover() ; error !=nil {
		fmt.Println("Seems Program does not finish correctly ")
	}      
 }()	
  
 

// PANIC and RECOVER    
// PANIC-- stops ordinary flow of func and begin to pani.When on func F , the F func stops execution..Any DEFFERED func's in there executes NORMALLY.F returns to its caller.To the caller F behaves like a call to panic   

// Here we are calling Open on a file want Opened (globalVar.go)
// Check error on filepath provided above 
if file, error :=os.Open("globalvare.go") ; error != nil {   

        // throws panic before including  defer func above
		panic("No possible to obatain the file!")
	}else{      

         defer func ()  {
			  fmt.Println("File is closed!") 
			  file.Close() 
		 }()                               
                               
		// if file is found we create an object(slice) to it and assign size  of file to be accessed (1254 byte)           
		 content := make([]byte, 1254)
		//  Read func read content provided 
		//  long--length of file to be read thus [:long]--meaning read content until long index(LAST)
		 long, _ :=file.Read(content)  
		//  string()--sequence of variabl-width xters where each n every xter is represented by one or more bytes using UTF-8 encoding 
		 contentArchived :=string(content[:long]) 
		 fmt.Println(contentArchived)
	}                        
    //  We are closing the opened file 
	
}   
    