package main

import "fmt"

type student struct {
    name string
	id   int
}

type myString string

func (s myString) count() int {
	i:=0
	for range s {
        i++
	}
	return i
}


func (s student) details() {
	fmt.Println("name:",s.name,"id:",s.id)
}


func (s *student) editDetails(newName string,newId int){
	s.name,s.id = newName,newId
}

type song struct {
	 name, singer, movie string

}


func (s song) print() string{
	return "song: "+ s.name + " singer: "+ s.singer +" movie: "+s.movie
}

func methodExamples(){

	s:= student{"nikhil",1}
	s.details()

	var name myString = "nikhil"
	fmt.Println(name.count())

	s.editDetails("mara nikhil",2001)
	s.details()

	a := &student{"nikhitha",2}
	a.details()
	a.editDetails("mara nikhitha",1111)
	a.details()


	song1:= song{"vasta ne venuka","harini","nani"}
	fmt.Println(song1.print())

}