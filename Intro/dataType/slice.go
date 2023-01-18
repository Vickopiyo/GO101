package main   
import "fmt"     

// Array has a fixed size whereas Slice is dynamiccally sized . flexible view into the elements of an array.
// []T---looks like array notation but with capital T at the end of bracket 
// Slice doesnt store any value.it just  describes a section of underlying array       

// numbers :=[4]int{1,2,3,4}--This is an Array.Fixed size   
// numbers :=[]int{1,2,3,4}--This is a slice.Dynamically sized   
 
func main()  {                 
      
     numbers :=[]int{1,2,3}       
    fmt.Println(numbers, len(numbers))   
    // Appending values to slice using append(arrayTobeAppended, valueToBeAppended)  
	numbers = append(numbers, 4)       
    numbers = append(numbers, 5)  
    fmt.Println(numbers, len(numbers))         
	// SubSlice --uses current values of the array despite reassigning happening after it    
	  subSlice :=numbers[:2]   
       numbers[0]= 100     
	  fmt.Println(subSlice)      
	  fmt.Println(numbers)     

	//More examples of Slices 
	  months :=[]string{"Jan", "Feb", "March", "April"}     

	  fmt.Printf("Length of months: %v cap: %v,  %p \n",  len(months), cap(months), months)      
	  
	//   Difference btw len() and cap()--len gives the total length of slice and can used to looping through while cap() total number of elements in an underlying array. 
  
}