package main     
  
import "fmt"
 
func main()  {

// How to declare array  
// bracket notation to indicate its of arrray type then number of items inside the bracket array and type of elements outside the bracket notation
// array declared with int no values  default to 0's in the array  
var numbers [5]int     

// assigning using 'indeces 
 numbers[0] = 20 
 numbers[1] = 30 
 numbers[2] = 40    
 numbers[3] = 50

  fmt.Println(numbers) 
  fmt.Println(numbers[3])        

//Another way of initializing array 
names :=[4]string{"VIck","Okoth", "Opiyo", "Mboya"}  
ages :=[3]int{2,4,5}
fmt.Println(ages)  
fmt.Println(names)   

// ELLIPSIS--used to pass varying number of arguments into a func(variadic--func that takes varying number of argument) or array 
colors :=[...]string{"Red","blue", "green"}  
// len(arr)--returns number of element in the array or any datatype passed into it.
    fmt.Println(colors, len(colors))  
	// Assigning using indeces.Always end with a comma because the array can pick more elements 
    schools :=[...]string{
		0: "Agoro",
		1: "Agosa",
		2: "Sare",
		3: "Alcatraz",
		4: "Yenga",
	}           
	fmt.Println(schools, len(schools))
	// subArrays--from index 0 to index 3.It does not  
	subSchool := schools[0:3]      
	fmt.Println(subSchool, len(subSchool))    
	// Passing of data in arrays by Value and reference
	//  BY VALUE--games  uses the values used intializing   
	
	games :=[...]string{"Football", "Hockey", "Rugby", "Basketball"}  
	// Value passed by Values 
	games1 := games  
	fmt.Printf("Original string is :  %v \n", games)  
	fmt.Printf("Array passed by Value: %v  \n", games1)   
	// BY REFERENCE --the amperstand (&) checks the current value of schools  and uses it 
      games2 := &games    
    
	// Reassigning schools to see diff in passing data using value and by reference   
       games[0]= "Futaaa"  
	   fmt.Printf("Original games  is : %v \n", games)   
	   fmt.Printf("Games 1 is : %v \n", games1)     
	//    Games 2 was passed by reference and the asterisk is used as pointer deference(takes value of variable ponted to it)
	   fmt.Printf("Games 2 is %v \n ", *games2)
     
	}   