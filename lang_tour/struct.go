package main

import "fmt" 

func structExamples(){

//define a struct

type movie struct {
	name string
	cast []string
	releaseYear int
}

emptyMovie:=movie{}
fmt.Println(emptyMovie,emptyMovie.name,emptyMovie.cast,emptyMovie.releaseYear)
ssmbNextMovie:= movie{"sarkaru vari pata",[]string{"mahesh","keerthi"},2022} 
fmt.Println(ssmbNextMovie)

ccNewMovie := movie{name:"cindrella",releaseYear:2021}
fmt.Println(ccNewMovie)

favMovie := struct {
	name string
	cast []string
	releaseYear int
}{name:"shershah",releaseYear:2021}
fmt.Println(favMovie.name)

p:=&favMovie
fmt.Println(p)

q:=&movie{"rrr",[]string{"ramcharan","ntr"},2021}
fmt.Println(q.name,(*q).releaseYear)

}