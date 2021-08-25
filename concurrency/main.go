package main

import (
	"fmt"
    "time"
)

func printNumbers(channel chan int) {
	for i:=0 ; i<=10;i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

    channel<-5

}

func main() {
	fmt.Println("concurrency with go")
     
	channel := make(chan int)
	go printNumbers(channel )
	lastNumber:= <-channel
	fmt.Println(lastNumber)

}