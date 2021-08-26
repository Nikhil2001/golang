package main

import "fmt"

const name ="nikhil mara"

const (
	language = "Go"
	year = 2021
)


const (	
	a  = iota
	b
	c
)

const (	
	x  = iota
	y  = iota
	z  = iota
)

const (	
	_  = iota
	i  
	j  = i + 1
	k
)


func constantExamples(){

	const age = 27

	fmt.Println(name,age)
	fmt.Println(language,year)
	fmt.Println(a,b,c)
	fmt.Println(x,y,z)
	fmt.Println(i,j,k)
	

}