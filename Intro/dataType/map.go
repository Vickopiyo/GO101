package main

import (
	"fmt"
)

func main() {

	// Maps map keys to values.And returns array with key-value pairs as eleements
	// Map initialization--datatype inserted in the bracket notation ..with both data types of key and value..Key in the bracket and Value in the outside bracket
	subjects := make(map[int]string)
	fmt.Println(subjects)
	// Nothing like indeces here..Map is more of an object in JS bt with bracket notation
	//fThe Numbers act like Keys--
	subjects[1] = "Math"
	subjects[2] = "Chem"
	subjects[3] = "Hist"
	subjects[4] = "Bio"
	fmt.Println(subjects)
	// Reassigning a value in a map
	subjects[3] = "History"
	fmt.Println(subjects)
	// Delete a key of a map .Accepts two args(map, keyToBedeleted)
	delete(subjects, 4)
	fmt.Println(subjects, len(subjects))
	//Map with value of slice
	students := make(map[string][]string)
	students["VICK"] = []string{"Chem", "Bio", "Math"}
	students["Eddy"] = []string{"Phyc", "Geo", "BS"}
	fmt.Println(students)  
	// Accessing Value(s)--s if the value is array/slice  of a key in map 
	fmt.Println(students["VICK"][2])    
   
}
