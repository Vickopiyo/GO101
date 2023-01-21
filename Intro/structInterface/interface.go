package main  
import "fmt"
              

type Animal  interface { 
	mover() string 
}

type Dog  struct {
	
} 
type Fish struct {
	
}    
type Bird struct {
	
}         

func (*Dog) mover() string{
	 return "I am dog and I move"
}     

func (*Fish) mover() string{
	return "I am fish  and I move"
}        

func (*Bird) mover() string{
	return "I am bird and I move"
}         
   

               

func  moverAnimal(animal Animal)  {     
	  fmt.Println(animal.mover())
}
           
       
func main()  {    
    //  INTERFACE---set oof method signatures 
	// Value of interface type can hold any value that implements those  methods.  
        
	hushpuppy :=Dog{}     
	moverAnimal(&hushpuppy)  

	tilapia :=Fish{}         
	moverAnimal(&tilapia)

	weaverbird :=Bird{}
	moverAnimal(&weaverbird)      
	
	
}