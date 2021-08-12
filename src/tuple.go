package main


import "fmt"

func ab()(int,int){
	return 3,4
}


func abc() (a int,b int,c string) {
	a,b,c=1,2,"go"
	return 
}

func exTuple() {
 a,b := 1,2 
fmt.Println(a,b)
fmt.Println(ab())
x,y,z := abc()
fmt.Println(x,y,z)


}