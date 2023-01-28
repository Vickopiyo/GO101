package main     
  
import "fmt"
 
func main()  {

// How to declare array  
// bracket notation to indicate its of array type then number of items inside the bracket array and type of elements outside the bracket notation
// Array declared with int no values  default to 0's in the array      

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
	subArraySchools := schools[3:]           

	fmt.Println("This is a sub array from index 3 ",subArraySchools)        
               
	fmt.Println(subSchool, len(subSchool))    
	// Passing of data in arrays by Value and reference
	//  BY VALUE--games  uses the values used intializing   
	games :=[...]string{"Football", "Hockey", "Rugby", "Basketball"}           
	     
	fmt.Printf("This a games pointer %p \n", &games)                          
                       
	// Value passed by Values                              
	games1 := games    
        
	fmt.Printf("Original string is :  %v  \n", games)               
    
	fmt.Printf("Array passed by Valu e: %p  \n", &games1)                  
	              
	// BY REFERENCE --the amperstand (&) refers the value of games2 to games.Ampersand points to that address.   

      games2 := &games             

	// Reassigning schools to see diff in passing data using value and by reference   
       games[0]= "Futaaa"                  

	   fmt.Printf("Original games  is : %v \n", games)           

	   fmt.Printf("Games 1 is : %v \n", games1)       
	//    Games2  value was passed by reference and the asterisk is used as POINTER DEFERENCE(takes value of variable pointed to it)----Asterisk show the passed pointed is variable with its value obtained by reference 

	   fmt.Printf("Games 2 is %v \n", *games2)             

	//    FILTERING Sub Arrays VALUES of Arrays using colons     
    //    The last index in the bracket notation must be greater than the first 
        fmt.Printf("sub Array from second to last %v, \n", games[1:])         
		// Last index in the bracket notation indicates where sub array ends.So it prints item[1], item[2], item[3]
		fmt.Printf("Sub Array from second item to third item: %v \n", games[1:3])  
		fmt.Printf("From start but ends at index 3 : %v \n", games[:3]) 
		fmt.Printf("Last item of array: %v \n", games[len(games)-1])            
	}                
	
	         