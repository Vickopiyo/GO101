package main  
import "fmt"  

// STRUCT begins with type then name of struct starting with a UPPERCASE 
// STRUCT--acts object in JS 

// STRUCT--PERSONA 

type  Person struct{  
    //   Attributes   
	name string  
	age int          
}        

//Methods 
//  *binds the intro func to  refer or an attribute of Persona  where p represent the Persona       


func (p*Person) intro(){

	fmt.Printf("Hi I am %s, %d of age \n", p.name, p.age) 
}   

//methods that receives  args 
func (p*Person)editName(newName string)  {
    p.name = newName
	fmt.Println(p.name) 
}        


// inheritance      
type Employee struct {
    //  employee inherits all attributes of Person
   Person   
   salary float64         

}

       
func main()  {    

	// person1 is a variable of Person.Accessing its properties----  person1.attribute= "reassigning too"
		person1 := Person{"Vick", 26}   

		person1.intro() 
		// INTRO called before reassigning   
		// person1.name= "VICTOR"     
		// passing a new name using a func   
		person1.editName("Vick97")
		// fmt.Println(person1)     
   // Accessing and reassigning attribute of a varible of struct 
       person1.age = 24  
   //  Strictness initiating attributes  
      person2 := Person{
		name: "Eddy",
        age: 34,
	  }         
	    person2.intro()
	//fmt.Println(person2)       
	
	// inheritance in action    
	em1 :=Employee{   

		salary: 6000,
	}      
	em1.name = "Pedro"  
	em1.age = 45    
	em1.intro()     
// Prints with double {{}}--employee is a struct which inherits attr for another struct.So the inner struct is the one inherited from 
	fmt.Println(em1)           



}       
         


