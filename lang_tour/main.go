package main

import "fmt"
import "math"
import ("time" //this is called factored or paranthesized import 
        "math/rand"
)

//about functions
func add(a int,b int) int {
	return a+b
}

func mul(a,b,c int) int {
	return a*b*c
}

func addAndSub(a,b int) (int,int) {
	return a+b,a-b
}

func swap(a1,b1 int)(a,b int){
    a,b = b1,a1
	return 
}

//variable at package level
var language string

func main() {

	fmt.Println("hello world")
	fmt.Print("hello ")
	fmt.Print("Nikhil\n")
	fmt.Println("now the time is :",time.Now())

	//any exported name should begin with capital letter
	// to be accessible outside its package (in the below example its Pi not pi)
	fmt.Println("The value of  PI is ",math.Pi) 
	fmt.Println("square root of 25 is ",math.Sqrt(25)) 
	fmt.Println("some random number",rand.Intn(111))

	fmt.Println("1+2 =",add(1,2))
	fmt.Println("1*2*3 =",mul(1,2,3))
	a,b := addAndSub(1,2)
	fmt.Println("1+2 =",a,"and 1-2 =",b)
	a,b = swap(a,b)
	fmt.Println("a is now",a,"and b is now =",b)

	
	fmt.Printf("%q\n",language)
	language = "goLang"
	fmt.Println(language)

	var x,y,z ="oh", false ,false
	fmt.Println(x,y,z)

	const x1,y1,z1 = "go", true ,1.5
	fmt.Println(x1,y1,z1)

	var i,j,k int
	fmt.Println(i,j,k)

	cplx := 1+5i
	fmt.Printf("type:%T,value %v\n",cplx,cplx)


	f := 1.0
	fmt.Printf("type:%T,value %v\n",f,f)

	//!TODO: do more examples with basic types and type casting
    var u =5
	var exInt float64 = float64(u)
	fmt.Println(exInt)
}