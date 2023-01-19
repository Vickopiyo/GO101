package main   

import "fmt" 
     
func main()  {     


	// For PERFOMANCE FOR is Better than for ...range 
	// For DRY for..range is better 
//    MAIN DIFF btw for and for...range 
	// The for ... range version creates a copy of the item in the slice you are iterating over. The other, does not and you must access it by reference.

	// looping an array of string using for loop       
      names :=[...]string{"Vick", "Okoth", "Opiyo"}     
   
	  for i :=0 ; i  < len(names); i++ { 

		    fmt.Println(names[i])
	  }
    //  LOOPING USING FOR--RANGE of foreach(Javascript )  
	// NB --when using for range...make sure index is used or leave an EMPTY dash(_).Go doesnt allow unused variables    
	// for range has three args index, element and array  
	 
	for _,  name := range names {
		    
		fmt.Println(name)
	}
//   for range with index printed or used  

for index, name := range names {   
	fmt.Println(index,name)
}
    
}