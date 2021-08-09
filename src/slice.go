package main

import "fmt"
func exSlice(){

	var intSlice []int
	fmt.Println(intSlice)

	intSlice2:= []int{0,1,2}
	fmt.Println(intSlice2)

	intSlice2 = append(intSlice2,4,4,6,7,9,10)
	fmt.Println(len(intSlice2))
	fmt.Println(cap(intSlice2))
	fmt.Println(intSlice2)

	intSlice2 = append(intSlice2,4)
	fmt.Println(len(intSlice2))
	fmt.Println(cap(intSlice2))
	fmt.Println(intSlice2)


	intSlice2 = append(intSlice2,6)
	fmt.Println(len(intSlice2))
	fmt.Println(cap(intSlice2))
	fmt.Println(intSlice2)

/* OUTPUT
*/

}