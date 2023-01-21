package main

import "fmt"

// ONE TO MANY RELATIONS


type Game  struct {  
   name string   
   sports []Sport
             
}    
         
type Sport struct {   
	name string      
	duration string    
	// Each sport belong to a GAME  
	game  Game 
	
}              

func main()  {
	
   // sport 1                
   football :=Sport{
	name: "Football",
	duration: "90 minuntes",
	
	}         

	//sport2  
	netball :=Sport{
		name: "Netball",
		duration: "60 minutes",    
	}      
    // Game1   
	ballgames :=Game{
       name: "ballgames", 
	   sports: []Sport{netball, football},
	}       
	
	//  stickgames :=Game{
	// 	name: "Stickgames",  
	// 	sports: []Sport{netball},      
	//  }

	 netball.game= ballgames    
	 football.game = ballgames                  

      fmt.Println(ballgames)      
	  fmt.Println(netball)
       
//   iterate through sports in Game printing each game 
for index, sport := range ballgames.sports {  
	fmt.Println(index, sport.duration)            
}   

}      


