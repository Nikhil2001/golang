package main
import "fmt"


//TODO add afterValue and valueBetweenNodes methods  
type node struct {
     value int
	 prev *node
	 next *node
	 
}


type doublyLinkedList struct {
     head *node

}


func (node *node) print(){
	fmt.Println(node.value,node.prev,node.next)
}

func (list *doublyLinkedList) AddToHead(value int){
	if list.head==nil{
      list.head= &node{}
      list.head.value=value
    }else {
	var node = node{}	
	node.value=value
	list.head.prev=&node
	node.next=list.head
    list.head = &node
   }
}


func (list *doublyLinkedList) AddToEnd(value int){
var node = &node{}
node.value =value
node1:=list.head
for ;node1.next!=nil;node1=node1.next{
}
node.prev=node1
node1.next=node
}

func exDoublyLinkedList() {

  var list doublyLinkedList
  list.AddToHead(1)
  list.head.print()

  list.AddToEnd(2)
  list.head.print()
  list.head.next.print()
 
  list.AddToHead(2)
  list.head.print()
  list.head.next.print()
  
  list.AddToHead(3)
  list.head.print()
  list.head.next.print()
 

}