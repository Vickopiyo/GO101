package main

import (
	"fmt"
	"strings"
)        

func reverseString(str string) string {
	// SPLIT takes two args  (string, separator)..separator can be spaces or nothing which is an empty string
	// Separator here is nothing so it output every element of the string even Spaces 
	strArr := strings.Split(str, "")
	//  Create an EMPTY array  where elements will be appended
	reverseArr := make([]string, 0)
	// Looping through arr(splitted string)  and
	for i := len(strArr) - 1; i >= 0; i-- { 
		// Loops from the last element of arrStr and pushes each element to reverseStr
		reverseArr = append(reverseArr, strArr[i])     
	}                 
	// fmt.Println(strArr)
	// fmt.Println(reverseArr)     
	// separator is NOTHING because we want a tight string no with spaces or some characters   
	return strings.Join(reverseArr, "")       
}    

func isPalindrome(name string) bool{

	// Strings package has strings which has a method toLOWER which does the thing..Goes through each element in the string
	name = strings.ToLower(name)
	// REPLACING some elements in a string..  it has three args (string, oldelement, newelement, numberofreplacementInstances)
	// REPLACING is STRICT the element to be replaced must EXACLTY match element in the string..IN CASES too(low and upp)
	//    name = strings.Replace(name, "Z", "S",2)
	//    fmt.Println(name)

	//Diff btw REPLACE AND REPLACEALL is the replace has 4 args whereas replaceall has 3 args..
	//REPLACEAll spaces    with nothing thus compactig the string 
	name = strings.ReplaceAll(name, " ", "")  

	// NB NB---  2nd format below executes FIRST Bcz reverseString apppears(compiled first ) first before isPalindrome
	//  fmt.Println(name)
     fmt.Println(reverseString(name))    
	 
	//  reversed string of the input 
	 reverseName := reverseString(name)     
//   is the name passed similar to its reverse(palindrome)
	 return name == reverseName 
	      
}
     
func main() {
	if isPalindrome("Luz azul")   {
		fmt.Println("Hell yeah! its a palindrome")
	} else{
		fmt.Println("OOH its not a palindrome")     
	}
	// reverseString("luz azul")
}    
