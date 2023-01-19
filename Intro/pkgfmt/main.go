package main

import "fmt"

func main() {
	//   Prints concatenates  two strings without space and outputs
	salamu := "Hujambo"
	greeting := "Shikamoo"
	fmt.Print(salamu)
	fmt.Print(greeting)
	name := "Victor"

	occupation := "Programmer"    

	age := 26
	fmt.Printf("My name is %s.I am %s and I am %d years  old ", name, occupation, age)     
	fmt.Printf("My name is %v.I am %v and I am %v years  old ", name, occupation, age)

	introduction := fmt.Sprintf("My name is %s.I am %s and I am %d years  old ", name, occupation, age)
      
	fmt.Println(introduction)         
                    
	fmt.Printf("My age is %T :", age)
      
	// Println--prints to standard output  but with a new  line at the end as a default   
	// default space added at the end of printing mutiple values and  variables 
	// Print-prints output in a line and concatenates with another print     
	// PRINTF----prints outputs but formattted accdording to description above ..Does not add new line by default    
	// Sprintf is similar to printf writes data into character Array whereas printf  writes data into stdout, standard output device
	// %s--string %d--digit  while %v--default format which is prefferable because variable my change in type
	// %T outputs type of value whether int, string, bool, e.t.c   
	// %p--pointer in memory
      
	 fmt.Print("What's your telephone number? ")
	//  Scanln--acts like prompt in jS..It intepretes like above it and takes the input and it is printed below with  no fmt before print
	// last print is in lower case.. amperstand is usually placed before the input varible   
         telephone := 701341969   
		// Refers to the last value if no value is inputted       
     fmt.Scanln(&telephone)             
	 print("My telephone number is :", telephone)                 
}     
                           
     
