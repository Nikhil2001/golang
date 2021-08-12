package main

import "fmt"


type set struct {
     elements map[int]bool

 } 

 func (set *set) New(){
	set.elements = make(map[int]bool)

 }

 func(set *set) ContainsElement(value int) bool {
    _,present := set.elements[value]
   return present

 }

func (set *set) AddElement(value int ){
	if !set.ContainsElement(value){
       set.elements[value]=true
	}
}

func (set1 *set) union(set2 *set) *set{
    var setUnion set
	setUnion.New()
	var newSet1,newSet2 *set =set1, set2
	if len(set1.elements) < len(set2.elements){
		newSet1 =set2
		newSet2 =set1
	}
	setUnion = *newSet1
	for key:= range newSet2.elements{
		if _,prs:= setUnion.elements[key]; !prs==true{
		setUnion.AddElement(key)
	   }
	} 
	 return &setUnion

}


func (set1 *set) intersection(set2 *set) *set{
    var setInter set
	setInter.New()
	var newSet1,newSet2 *set =set1, set2
	if len(set1.elements) < len(set2.elements){
		newSet1 =set2
		newSet2 =set1
	}
	for key:= range newSet2.elements{
		if _,prs:= newSet1.elements[key]; prs==true{
			setInter.AddElement(key)
	   }
	} 
	 return &setInter

}


func (set *set) deleteElement(value int ){
       delete(set.elements,value)
}

func exSets(){
var set1,set2 set
set1.New()
set2.New()
set1.AddElement(11)
set1.AddElement(12)
set1.AddElement(13)
set1.AddElement(14)

set2.AddElement(21)
set2.AddElement(22)
set2.AddElement(23)
set2.AddElement(24)

fmt.Println(set1.union(&set2))

set1.New()
set2.New()
set1.AddElement(1)
set1.AddElement(2)
set1.AddElement(3)
set1.AddElement(4)

set2.AddElement(4)
set2.AddElement(5)
set2.AddElement(7)
set2.AddElement(1)

fmt.Println(set1.union(&set2))


set1.New()
set2.New()
set1.AddElement(1)
set1.AddElement(2)
set1.AddElement(3)
set1.AddElement(4)

set2.AddElement(1)
set2.AddElement(2)
set2.AddElement(3)
set2.AddElement(7)

fmt.Println(set1.union(&set2))


set1.New()
set2.New()
set1.AddElement(11)
set1.AddElement(12)
set1.AddElement(13)
set1.AddElement(14)

set2.AddElement(21)
set2.AddElement(22)
set2.AddElement(23)
set2.AddElement(24)

fmt.Println(set1.intersection(&set2))

set1.New()
set2.New()
set1.AddElement(1)
set1.AddElement(2)
set1.AddElement(3)
set1.AddElement(4)

set2.AddElement(4)
set2.AddElement(5)
set2.AddElement(7)
set2.AddElement(1)

fmt.Println(set1.intersection(&set2))


set1.New()
set2.New()
set1.AddElement(1)
set1.AddElement(2)
set1.AddElement(3)
set1.AddElement(4)

set2.AddElement(1)
set2.AddElement(2)
set2.AddElement(3)
set2.AddElement(7)

fmt.Println(set1.intersection(&set2))


/*OUTPUT
nikhil@nikhil:~/golang$ go run src/main.go src/sets.go 
Hello World!
&{map[11:true 12:true 13:true 14:true 21:true 22:true 23:true 24:true]}
&{map[1:true 2:true 3:true 4:true 5:true 7:true]}
&{map[1:true 2:true 3:true 4:true 7:true]}
&{map[]}
&{map[1:true 4:true]}
&{map[1:true 2:true 3:true]}
*/


}