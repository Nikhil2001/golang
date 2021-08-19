package main

import "fmt"


func rangeExamples(){

	twoTable := [10]int{2,4,6,8,10,12,14,16,18,20}

	for i,v:= range twoTable {
		fmt.Printf("2 x %d = %d \n",i+1,v)
	}

	for i,_:= range twoTable {
		fmt.Println("index ",i) //-------
	}                           //      |
                                //      |-------> both are same
	for i := range twoTable {   //      |  
		fmt.Println("index ",i) //-------
	}

    studentIds:= map[string]int{"nikhil":1,"nikhitha":2}	
	fmt.Println(len(studentIds))
	for key := range studentIds {
		fmt.Println(key)
	}

  for i, v := range "nikhil" {
	  fmt.Println(i,v)
  }

	
}