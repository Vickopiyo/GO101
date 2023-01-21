package main      
import "fmt"
   


type Staff struct{   
	name string    
	age int         
	salary int  
	isPermanent bool
	role string
        
}
   
type Doctor struct{          

    // Here Staff is an embedded struct thus comes with poromoted fields like name,age etc
	Staff 
	isAdmin bool  

}    

type Nurse struct{   

    Staff        
	hasOvertime bool

}    

type Intern struct {  

	Nurse         
	duration string   

}           

// In Go, these fields from embedded structs are called promoted fields.

func main()  {       

	// Regular staff 
	cleaner := Staff{
		name: "Vick",   
		age: 24,   
		salary: 3000, 
		role: "Cleaner", 
		isPermanent: false,
	}        
	fmt.Println(cleaner)    

	// DOCTOR 
	daktari :=Doctor{
        isAdmin: true,      
        Staff: Staff{
			name: "Dr.Opiyo",
		    age: 26,  
			salary: 40000,
			role: "Surgeon",
			isPermanent:  true,
		},

	}       
	fmt.Println(daktari)     
	// NURSE   
	nurse := Nurse{   
		Staff: Staff{
			name: "Lilian",  
			age: 23,
			salary: 20000,
			role: "Threatre Asisstant",
			isPermanent: false,    
		},
		         
		hasOvertime: true, 
	}  
	fmt.Println(nurse)   

//    INTERN  

intern := Intern{

	duration: "6 Months",   
	Nurse: Nurse{
		Staff: Staff{
			name: "Julian",
			age: 23,
			salary: 15000, 
			isPermanent: false,
			role: "Surgeon Intern",      
		}, 
		hasOvertime: false,  
	},
}     

fmt.Println(intern)

}