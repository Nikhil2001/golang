package main

import (
	    "fmt"
        "time" 
	)

func isEven(x int ) bool {
	return x%2 == 0
}

func getShape(x,y int ) (shape string) {
	if x==y {
		shape= "square"	
	}else {
		shape= "rectangle"	
	}
	return
}

func switchExamples() {

	num:=6

	if n:=isEven(num); n {
		fmt.Printf("%d is even  \n",num)
	}
   
	fmt.Printf("it's a %s \n",getShape(2,2)) 
	fmt.Printf("it's a %s \n",getShape(4,6)) 
	if shape:=getShape(4,5); shape=="square"{
		fmt.Println("det är en kvadrat")
	} else {
		fmt.Printf("det är en %s \n",shape)
	}

	if day:=time.Now().Weekday(); day==time.Monday {
		fmt.Println("its Monday")
	} else if day==time.Tuesday {
     	fmt.Println("its Tuesday")
	} else{
		fmt.Println("its **",day)
	}
	
	switch k:=2; k {
	case 1 :
		    fmt.Printf("k is %d\n",k)
	case 2,4 :
	    	fmt.Printf("k is 2 or 4\n")
			fallthrough
	case 3 :
		   fmt.Printf("k is %d\n",k)
	default:
		  fmt.Println("k is not 1 2 3")
	}
	switch  {
	case true :
		fmt.Println("case is  true")
	case false :
	    fmt.Println("case id false ")
	}

	 
}