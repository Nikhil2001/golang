package main

import "fmt"

func add(a int,b int) int {
	return a+b
}

func mul(a,b,c int) int {
	return a*b*c
}

func addAndMul(a,b,c int) (int,int) {
	return a+b+c,a*b*c
}

func addAndSub(a,b int) (add ,sub int) {
	add,sub = a+b,a-b
	return 
}

func addAndMulMany(nums ...int) (add , mul int){
    mul=1
	for _,num := range nums {
		add+=num
		mul*=num
	}
	return

}

func anyFunction(fn func(int, int)int) int {
	return fn(11,12)
}


func closure () func()int {
	i:=0
	return func() int {
		i++
		return i
	}
}

func functionExamples(){
	fmt.Printf("1+2=%d\n",add(1,2))
	fmt.Printf("1*2*3=%d\n",mul(1,2,3))
	a,m:=addAndMul(1,2,3)
	fmt.Printf("1+2+3=%d,1*2*3=%d\n",a,m)
	add,sub:=addAndSub(1,2)
	fmt.Printf("1+2=%d,1-2=%d\n",add,sub)

	a,m=addAndMulMany(1,2,5)
	fmt.Printf("add=%d,mul=%d\n",a,m)

	var addFn  = func(a,b int) int {return a+b}
	fmt.Println(addFn(1,2),anyFunction(addFn))
	//closure1:=closure()()
	fmt.Println(closure()())
	fmt.Println(closure()())

	closure1:=closure()
	fmt.Println(closure1())
	fmt.Println(closure1())
	
	
}


