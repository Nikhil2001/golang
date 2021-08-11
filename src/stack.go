package main

import "fmt"


type stack struct {
 elements []int

}

func (stack *stack) Push(value int){
    stack.elements =append(stack.elements,value)


}

func (stack *stack) Pop() *int {
	len:=stack.Len()
	if len>0 {
	lastElement :=stack.elements[len-1]
    stack.elements = stack.elements[0:len-1]
	return &lastElement
	}
	return nil

}


func (stack *stack) Len() int {
    return len(stack.elements)


}

func (stack *stack) Print(){
    for _,value:=range stack.elements{
		fmt.Print(value," ")
	}
	fmt.Println()

}

func exStack(){

	var stack1 stack;
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	fmt.Println("len:",stack1.Len())
	stack1.Print()
	fmt.Println("element deleted ",*stack1.Pop())
	fmt.Println("len:",stack1.Len())
	stack1.Print()

	fmt.Println("element deleted ",*stack1.Pop())
	stack1.Print()

	fmt.Println("element deleted ",*stack1.Pop())
	fmt.Println("len:",stack1.Len())
	stack1.Print()

	stack1.Pop()
	fmt.Println("len:",stack1.Len())

/* OUTPUT  
nikhil@nikhil:~/golang$ go run src/main.go src/stack.go 
Hello World!
len: 3
1 2 3 
element deleted  3
len: 2
1 2 
element deleted  2
1 
element deleted  1
len: 0

len: 0

*/
	

	


}