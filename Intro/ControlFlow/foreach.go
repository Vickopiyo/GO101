package main   

import "fmt" 

func main()  {  

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