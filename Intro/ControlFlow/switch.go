package main

import "fmt"

func main() {
	// SWITCH statements
	// A program to check if a letter is a vowel
	var vowel string
	fmt.Print("Please input a vowel?")
	fmt.Scanln(&vowel)


	// Code runs but not DRY--vowel keeps repeating itself 
	// switch {     
	// case vowel == "A" || vowel == "a":
	// 	fmt.Printf("%s is  a vowel!", vowel);    
	// case vowel == "E" || vowel == "e":
	// 	fmt.Printf("%s is  a vowel!", vowel);   
	// case vowel == "I" || vowel == "i":
	// 	fmt.Printf("%s is  a vowel!", vowel); 
	// case vowel == "O" || vowel == "o":
	// 	fmt.Printf("%s is  a vowel!", vowel); 
	// case vowel == "U" || vowel == "u":
	// 	fmt.Printf("%s is  a vowel!", vowel);  
	// default:
	// 	fmt.Printf("%s is not a vowel",vowel)
	// }
    
	// DRY---switch statement    
	// switch takes in conditon/variable as arguement and checks which case matches it and executes that block of code 
	switch vowel {
	case "A" , "a", "E", "e","I", "i","O","o","U","u":   
		fmt.Printf("%s is a vowel!", vowel)
	default:
		fmt.Printf("%s is NOT a vowel", vowel)	     
	}
    
}

