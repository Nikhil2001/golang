package main

import "fmt"


func exMap(){

var map1 = make(map[int]string)
fmt.Println(map1)

map2 := map[int]string{1:"go",2:"c++"}
fmt.Println(map2)

fmt.Println(map2[2])

map2[2]="rust"
fmt.Println(map2)

value,present:= map2[2]
fmt.Println(value,present)
delete(map2,2)

value,present= map2[2]
fmt.Println(value,present)

map3:= map[int]bool{1:false,2:true}
delete(map3,2)

value1,present1:= map3[2]
fmt.Println(value1,present1)

map2[2]="erlang"
map2[3]="swift"
fmt.Println(map2)

for key,value:= range map2 {
	fmt.Println(key,value)
}
/*OUTPUT:

nikhil@nikhil:~/golang$ go run src/main.go src/map.go 
Hello World!
map[]
map[1:go 2:c++]
c++
map[1:go 2:rust]
rust true
 false
false false
map[1:go 2:erlang 3:swift]
1 go
2 erlang
3 swift
nikhil@nikhil:~/golang$ 

*/


}