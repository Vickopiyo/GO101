package main

import "fmt"


type User struct {
	name   string
	email  string
	active bool
}

type Student struct {
	user User
	code string
}  

// COURSE can have MANY VIDEOS thus ONE TO MANY 

type Course  struct {      

	title string     
   // videos field is an empty slice with acepts elemnts which are instances of Video(struct)
	videos [] Video
}                      

// VIDEO BELONGS TO COURSE   one to one    

type Video  struct {
	title string    
	// Establishing rship btw video and course.Video id From Course 
	course Course         
}

func main() {
	
	// RELATION OF ONE TO ONE 
	// User1
	vick := User{
		name:   "vick",
		email:  "vick@gmail.com",
		active: true,
	}

	// user2
	okoth := User{
		name:   "okoth",
		email:  "okoth@gmail.com",
		active: true,
	}
	// Student  
	vickStudent := Student{
		user: vick,
		code: "9033",
	}            
     
	fmt.Println(okoth)      
	// Accessing attributes of Original User       
	fmt.Println(vickStudent.user.name)           

    //  RELATION ONE TO MANY    
	videos1 :=Video{title: "O1-intro GO"}          
	videos2 :=Video{title: "02-Variables in GO"}            
	
	goLang :=Course{                    
		title: "GOLANG 101" ,         
		//  Adding videos to course 
		videos: []Video{videos1, videos2}, 

		// Assigning videos to courses       
		}   
       videos1.course = goLang      
	   videos2.course = goLang          
       fmt.Println(goLang)          
       fmt.Println(videos1.course)    

}
          
   

