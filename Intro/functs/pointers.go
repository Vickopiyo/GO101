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



  
func modifiedVariable(name *string)  { 
	*name = "Modified Variable"             
	fmt.Printf("%p \n", name)              
}        

            
func main()  {
	// POINTERS   
  name := "Main variable "   
  fmt.Printf("%p \n", &name)         
  fmt.Println("It is a ", name)      
  //  Here a modified variable points to its orignal variable due to its (name *string) 
  modifiedVariable(&name)     
  
//   modifiedVar REASSIGNS(remeber ?? *name="modified variable"??) name variable to its original variable
  fmt.Println("Which var is this ? ", name)       
	
}     