package main


import (
	  "fmt"
     "math" 
	)

type shape interface {
	area() float64
	perimeter() float64
}


type circle struct {
	radius float64
}


type square struct {
	length float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return   2 * math.Pi * c.radius 
}

func (s *square) area() float64 {
	return  s.length * s.length 
}

func (s *square) perimeter() float64  {
	return   4 * s.length 
}


func shapeParams(s shape){
	fmt.Printf("%v %T\n",s,s)
	fmt.Println(s.area(),s.perimeter())
}


func interfaceExamples(){
	shapeParams(circle{7})
	shapeParams(&square{2})
	
}
