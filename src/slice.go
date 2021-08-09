package main

import "fmt"
func exSlice(){

	var intSlice []int
	fmt.Println(intSlice)

	intSlice2:= []int{0,1}
	fmt.Println(intSlice2)

	intSlice2 = append(intSlice2,4,5,6)
	fmt.Println(intSlice2,len(intSlice2), cap(intSlice2))
	intSlice2 = append(intSlice2,7)
	fmt.Println(intSlice2,len(intSlice2), cap(intSlice2))

	intSlice2 = append(intSlice2,8)
	fmt.Println(intSlice2,len(intSlice2), cap(intSlice2))

	intSlice3 := make([]int,3)
	fmt.Println(intSlice3,len(intSlice3), cap(intSlice3))

	intSlice3 = append(intSlice3,8)
	fmt.Println(intSlice3,len(intSlice3), cap(intSlice3))

	var intSlice4 []int
	copy(intSlice4,intSlice3)
	fmt.Println(intSlice4,len(intSlice4), cap(intSlice4))

    intSlice5 := make([]int,4)
	copy(intSlice5,intSlice3)
	fmt.Println(intSlice5,len(intSlice5), cap(intSlice5))

    stringSlice := []string{"i","am","a","go","programmer","."}
	fmt.Println(stringSlice,len(stringSlice), cap(stringSlice))


	stringSlice1 := stringSlice[0:]
	fmt.Println(stringSlice1)

	stringSlice2 := stringSlice[:5]
	fmt.Println(stringSlice2)

	stringSlice3 := stringSlice[0:5]
	fmt.Println(stringSlice3)

	var stringSlice4 []string
	stringSlice4 = stringSlice[0:6]
	fmt.Println(stringSlice4)

	intSlice2d:= [][]int{{1},{1,2},{1,2,3}}
	fmt.Println(intSlice2d,len(intSlice2d), cap(intSlice2d))
	fmt.Println(intSlice2d[2])

	intSlice2d_1:= make([][]int,3)
	copy(intSlice2d_1,intSlice2d)
	intSlice2d_1 = append(intSlice2d_1,[]int{1,2,3,4})
	fmt.Println(intSlice2d_1,len(intSlice2d_1), cap(intSlice2d_1)) 

	var intSlice2d_2 [][]int

	for i:=0;i<3;i++ {
		intSlice2d:=make([]int,i+1)
		for  j:=0;j<len(intSlice2d);j++{
			intSlice2d[j]=j
		
		}
		intSlice2d_2=append(intSlice2d_2,intSlice2d)
	}

	fmt.Println(intSlice2d_2,len(intSlice2d_2), cap(intSlice2d_2))


/* OUTPUT

nikhil@nikhil:~/golang$ go run src/slice.go src/main.go 
Hello World!
[]
[0 1]
[0 1 4 5 6] 5 6
[0 1 4 5 6 7] 6 6
[0 1 4 5 6 7 8] 7 12
[0 0 0] 3 3
[0 0 0 8] 4 6
[] 0 0
[0 0 0 8] 4 4
[i am a go programmer .] 6 6
[i am a go programmer .]
[i am a go programmer]
[i am a go programmer]
[i am a go programmer .]
[[1] [1 2] [1 2 3]] 3 3
[1 2 3]
[[1] [1 2] [1 2 3] [1 2 3 4]] 4 6
[[0] [0 1] [0 1 2]] 3 4

*/

}