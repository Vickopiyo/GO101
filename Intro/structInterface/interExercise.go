package main

import (
	"fmt"
	"math"
)

// interface bundles up   related functions for structs  together in one func (measurement)--
type Geometric  interface {     

	area()  float64 
	perimeter()   float64
	  
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c *Circle) area() float64 {
	return  math.Pi*(c.radius* c.radius)

}      

func (p*Circle) perimeter() float64{    
  
	return 2 * math.Pi* p.radius 
      
}           

func (r*Rectangle) area() float64{    
	return r.height * r.width   

}       

func (r*Rectangle)  perimeter() float64 {
	   
	return 2*(r.height + r.width)     
}        

// receives a  struct which has func related enlist in the interface       
func measurement( g Geometric)  {      

	  fmt.Println(g) 
	  fmt.Println(g.area())    
	  fmt.Println(g.perimeter())            
}

func main() {          
   
	circle1 := Circle{radius: 23 }      
	fmt.Println(circle1.area())
    fmt.Println(circle1.perimeter())        
	rec1 := Rectangle{width: 20, height: 15}
    fmt.Println(rec1.area())  
	fmt.Println(rec1.perimeter())          

	measurement(&circle1)        
	measurement(&rec1)         
    
}   

