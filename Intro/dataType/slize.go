package main

import "fmt"

// Array--indexed colection of certain size with values of same type
// Array--are passed by values(making a copy of the original)
// Arrray CANNOT  be resized i.e var arr [3]int --means array can hold only 3 items

// Slice--a data structure describing a piece of an array with three properties namely: pointer,size and capacity
// capacity--length of slice and also inidcates the maximum length the slice can take(as it gows)

func PrintSlize(slize []int) {

	fmt.Printf("Length=%d Cap=%d , %v \n", len(slize), cap(slize), slize)
}

func main() {

	// Slices must in range
	numbers := []int{1, 2, 3, 4, 5, 6}
	//  outputs len-6 cap-6
	PrintSlize(numbers)
	//     Slice the slice to give it 0 lenghth
	sliceNumbers := numbers[:0]
	PrintSlize(sliceNumbers)
	//  Extend length of slice
	sliceNumbers = sliceNumbers[:4]
	sliceNumbers[2] = 344
	PrintSlize(sliceNumbers)
	// Drop its first two values...Cutting slice from a certain number of array reduces the capacity.Seen below
	sliceNumbers = sliceNumbers[2:]
	PrintSlize(sliceNumbers)
	PrintSlize(numbers)
	//  append is like the push() in JS only that it takes two arg..1.arr 2. itemToBepushed
	// How slice allocates doubles capacity if length is greater than capacity
    
	holidays := []string{"X-mas", "Labour", "Madaraka"}

	fmt.Printf("Famous holidays in Kenya are %d and capacity %d  while year(pointer)  %p \n", len(holidays), cap(holidays), holidays)
//   capacity becomes 6 because array length is greater than capacity thus more space for the Uhuru holiday to fit in.
// This changes the pointer or rather the addreess of the slice too
	holidays = append(holidays, "Uhuru")
	fmt.Printf("Famous holidays in Kenya are %d and capacity %d  while year(pointer)  %p \n", len(holidays), cap(holidays), holidays)

	var s []int
	for i := 0; i < 10; i++ {
		//outputs capacity changes each time length exceeds it..It actually doubles the current  length to form a new capacity thus cap 1,1,2,2,3,4,4,8,8,16,16
		// Address also keeps changing thus slice is highly mutable
		fmt.Printf("length: %d, capacity: %d, address: %p\n", len(s), cap(s), s)
		s = append(s, i)

		//  slices using make(arr,len, cap)

	}

}
