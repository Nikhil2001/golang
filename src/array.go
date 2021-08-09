package main

import "fmt"
func exArray(){

	var intArray [5]int
	fmt.Println(intArray)

	lenIntArray := len(intArray)

	for i:=0;i<lenIntArray;i++{
		intArray[i]=i;
	}
	fmt.Println(intArray)

	intArray[0]=1;
	fmt.Println(intArray)

	intArray2 := [5]int{11,12,13,14,15}
	for i,j:= range intArray2 {
		fmt.Println(i,j)
	}

	intArray2d := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(intArray2d)

	var intArray2d_2 [3][3]int
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			intArray2d_2[i][j]=j;
		}
	}
	fmt.Println(intArray2d_2);

	for index,value := range intArray2d_2 {
		fmt.Println(index,value)
	}


/* OUTPUT

nikhil@nikhil:~/golang$ go build src/array.go src/main.go 
nikhil@nikhil:~/golang$ ./array 
Hello World!
[0 0 0 0 0]
[0 1 2 3 4]
[1 1 2 3 4]
0 11
1 12
2 13
3 14
4 15
[[1 2 3] [4 5 6]]
[[0 1 2] [0 1 2] [0 1 2]]
0 [0 1 2]
1 [0 1 2]
2 [0 1 2]
nikhil@nikhil:~/golang$ 

*/

}