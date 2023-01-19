package main

import "fmt"

func main() {
	// Here we are going indepth on operators

	a := 2
	// Addition
	// a = a + 2
	// Concise way of using operators
	a += 2
	fmt.Println(a)
	// Subtracting
	a -= 2
	fmt.Println(a)
	// Multplying
	a *= 4
	fmt.Println(a)
	// Division
	a /= 2
	fmt.Println(a)
	// Modulus
	a %= 2
	fmt.Println(a)

	// INCREMENTING
	b := 0
	b++
	b++
	b++

	fmt.Println(b)
	// DECREMENTINIG
	b--
	b--
	fmt.Println(b)

}
