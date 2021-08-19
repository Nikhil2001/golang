package main

import "fmt"

func sliceExamples(){

	var slice1 []int
	fmt.Println(slice1,len(slice1),cap(slice1))

	slice2:=[]int{1,2,3,4,5}
	fmt.Println(slice2,len(slice2),cap(slice2))

	slice2 = append(slice2,6)
	fmt.Println(slice2,len(slice2),cap(slice2))

	slice3 := append(slice2,7)
    fmt.Println(slice3,cap(slice3))
	slice2[0]=11
	fmt.Println(slice2,slice3)
	slice2[0]=1
	fmt.Println(slice2[:6],slice2[:],slice2[0:],slice2[1:4])
	slice4:=make([]int,1)
	fmt.Println(slice4)
	copy(slice4,slice2)
	fmt.Println(slice4)

	slice5:=[]int{1,2,3,4,5}
	fmt.Println(slice5,len(slice5),cap(slice5))
	slice5 =append(slice5,6)
	fmt.Println(slice5,len(slice5),cap(slice5))
	slice6 :=slice5[5]
	fmt.Println(slice6)
}