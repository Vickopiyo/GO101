package main  
import "fmt"      


// Even though the two NAME variable are similar in name.They have differently scoped.DIfferent memory address.
// modifiedVaribale()  is called in main()  so it check for the variable in main block for a variable named name  
// modifiedVariable(name)--without & passes the value of name but creates memory bcz it was referenced( use of *)
// Even though it outputs same thing with &name but it is memory expensive   
// &name is used to create reference but * is used to bind it  especally when passing the reference in args  thus passing 
// modifiedVariable(name *string)---this  * binds name to scope it is  found..Any reasinging of name comes *name= "reasssigned "
// How to call modifiedVriable(&name)---references name in ITS SCOPE thus using same memory address.
// NUTSHELL-- &--REFERENCES *--BINDS (pointer deference)    
// Ampersand  -- &name makes a variable from a different scope available in the scope its in.  


func modifiedVariable(name *string)  { 
	*name = "Modified Variable"             
	fmt.Printf("%p \n", name)              
}        

func main()  {
	// POINTERS--All pointer spoint to same address 
  name := "Main variable "   
//   checks where it is referenced from!
  fmt.Printf("%p \n", &name)      
  // checks where where name is the that scope
  fmt.Println("It is a ", name)      
  // Here a modified variable points to its orignal variable due to its (name *string) 
  modifiedVariable(&name)     
//modifiedVariable  REASSIGNS(remember ?? *name="modified variable"??) name variable to its original variable
  fmt.Println("Which var is this ? ",name)   
  
  
	                                              
}     