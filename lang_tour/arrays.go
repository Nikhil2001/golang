package main

import "fmt"

func arrayExamples(){
	var array1 [5]int
	fmt.Println(array1,len(array1))
	array1[0]=11
	fmt.Println(array1)

	array2,array3:=[3]int{},[4]int{1,2,3,4}
	fmt.Println(array2,array3)
}