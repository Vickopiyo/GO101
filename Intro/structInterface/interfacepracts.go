package main

import (
	"fmt"
	"math"
)       


type Geometric  interface {      

	area() float64  
	perimeter()  float64
	              
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (side *Square) area() float64 {

	return side.side * side.side

}

func (side *Square) perimeter() float64 {     

	return side.side * 4

}                       
func (rd *Circle) area() float64 {

	return math.Pi*rd.radius*rd.radius
}
func (rd*Circle) perimeter() float64{    

	return 2*math.Pi*rd.radius
	       
}              
               
func measurement(g Geometric)  {
	       
     fmt.Println(g)         
	 fmt.Println(g.area())     
    fmt.Println(g.perimeter())
}
             
func main()  {          
   
       sq1 :=Square{
		side: 34 ,      
	   }      
	   
	   cr1 :=Circle{

		radius: 21,
	   }    

       measurement(&cr1)       
	   measurement(&sq1)               

}