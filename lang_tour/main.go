package main

import "fmt"
import "math"
import ("time" //this is called factored or paranthesized import 
        "math/rand"
		"runtime"
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

	//!TODO: do more examples with basic types,type casting
    var u =5
	var exInt float64 = float64(u)
	fmt.Println(exInt)

	//Loops
	for i:=0; i < 10;i++ {
		fmt.Println("for loop iteration:",i)
	}
	i =0
	for  i < 10 { //while like loop, without bool expression, its infinite loop
		fmt.Println("for loop iteration:",i)
		i++
	}
	i =0
	for  { //while like loop, without bool expression, its infinite loop
		fmt.Println("infinite loop iteration:",i)
		i++
		if i>3 {
			break
		}
	}

	if k:= true; k { //k can be decalred as variable with scope with in if 
		fmt.Println(!k)
	}

	if k:= false; k { //k can be decalred as variable with scope with in if 
		fmt.Println("in if, k is :",k)
	} else {
		fmt.Println("in else, k is :",k) //k also can be used in else block as well
	}
 fmt.Println("operating system:",runtime.GOOS)
 fmt.Println("Today is:",time.Tuesday==time.Now().Weekday())

 switch day:=time.Now().Weekday(); day {
 case time.Monday :
	fmt.Println("today is",day)
 case time.Tuesday :
	fmt.Println("today is",day)
  default :
	fmt.Println("today is",day)
 }
 
}