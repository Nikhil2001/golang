package main

import "fmt"


func pointerExamples() {

	var p *int
	if p == nil {
		fmt.Println("uninitialized pointer is",p)
	}

	i:=5;
	p=&i
	fmt.Println("intialized pointer",p)

	fmt.Println("access value through pointer",*p)
    *p =50
	fmt.Println("chnage value through pointer",*p)


}