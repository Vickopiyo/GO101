package main  
import "fmt"

func main()  {   

	//  6 Basic comaprison operators in GO'
	    //  > greater than
		//  < less than
		//  >=  greater than or equal to
		//  <=  less than or equal to 
		//  ==  equal to 
		//  !=  not equal to 


	    a := 4 
		b := 5 

		var r bool 
		// Comparison can ONLY be done to values of the same type 
		// inetgers,floats,booleans, complex numbers,pointers,string are COMPARABLE 
        r = a==b 
		fmt.Printf("Is %d equal to %d ? %t \n",a,b,r)   
		r =  a != b    
		fmt.Printf("Is %d equal to %d ? %t \n",a,b,r)       
		r =  a <= b     
		fmt.Printf("Is %d less than or equal to %d ? %t  \n",a,b,r)     
		r =  a >= b     
		fmt.Printf("Is %d greater than or equal to %d ? %t",a,b,r)       
      

}