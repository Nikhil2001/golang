package main

import "fmt"


func loopExamples() {

	 i := 0

	 for i<10 {
		 fmt.Print(i," ")
		 i++
	 }
	 fmt.Println()


	 for j:=9;j>=0;j-- {
		fmt.Print(j," ")
		
	}
	fmt.Println()

	for j:=9;j>=0; {
		fmt.Print(j," ")
		j--
	}
	fmt.Println()

	//infinite loop

	for {//infinite loop
	
		fmt.Println("breaking infinite loop")
		break;
		
	}

	for v:=0;v<10;v++{
		if v ==5 {
			continue;	
		}	
		fmt.Print(v," ")
	}
	fmt.Println()
}