//

package main

import "fmt"     
  

// TASK struct 
type Task struct{
	name string  
	desc string 
	completed bool      
}      

// Check task status 
func (t*Task) taskStatus()  {
	
	fmt.Printf("%s , %s complete? %t", t.name, t.desc, t.completed)
}
// Empty Structconatins task slice for updating tasks
type TaskList  struct {   
// an empty slice which hold data of type task(struct)   
	tasks []*Task                        
}        

// ADDING  task to the tasks in TaskList srruct 
func (tl*TaskList) appendTask(t*Task)  { 
	// Appending task append(arrayToReceiveAppends, itemTObeAppended )
	// itemTObeAppended CAN BE ITEM IN ANOTHER SO REST/ELLIPSIS WORKS like seen in removeTask 
  tl.tasks=append(tl.tasks, t)

}         

// REMOVING a task of specific INDEX 
func (tl*TaskList) removeTask(index int)  {  
	// copying elements fromindex + 1 postion to the array 
	tl.tasks=append(tl.tasks[:index], tl.tasks[:index + 1]... )   
	
}



func main()  {
	       
      task1 := Task{ 
		name: "GoLang", 
		desc: "Zero To Master GoLang", 
		completed: false,
	  }        
	  task2 :=Task{
		name: "React",
		desc: "Zero To Master React",
		completed: false,    
	  }      
	  
	  task3 :=Task{
		name: "Ruby",  
		desc: "Zero To Master Ruby", 
		completed: false,    
	  }    
	  

    // task1.taskStatus()  

	tasklist := TaskList{}   
   // Appendng Task 
	tasklist.appendTask(&task1)   
	tasklist.appendTask(&task2) 
	tasklist.appendTask(&task3 )        

	// fmt.Println(tasklist)        

	// iterate through  Tasklists   

	for index, task := range tasklist.tasks {   

		// Outputing Task Only comes & before it just to show it is a COPY of task(variable) in for range  
		fmt.Println(index, task.name)  
	}

   
}                