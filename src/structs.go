package main

import "fmt"

type person struct{
   weight int
   height int
}

func (person1 *person ) setPerson(h,w int){
	person1.weight,person1.height = w,h
}

func exStruct(){

	var person1 person
	fmt.Println(person1)

	person1.setPerson(182,70)
	fmt.Println(person1)

	person2 := person{30,70}
	fmt.Println(person2)

	person3 := person{height:30,weight:70}
	fmt.Println(person3)

	person4 := person{height:30}
	fmt.Println(person4)

/*OUTPUT

nikhil@nikhil:~/golang$ go run src/main.go src/structs.go 
Hello World!
{0 0}
{70 182}
{30 70}
{70 30}
{0 30}

*/

}