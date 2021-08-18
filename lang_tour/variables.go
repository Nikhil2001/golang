package main

import "fmt"

func variableExamples(){

var a,b int
fmt.Println(a,b) //prints 0,0 by default ints are 0s

var language string ="go"
var name,age = "nikhil",27
car,rocket := "mercedes", "falcon"

fmt.Println(language,name,age,car,rocket)

}