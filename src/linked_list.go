package main

import "fmt"


type node struct{
	value int
	next *node 
}

type linkedList struct{
	head *node
}


func (lList *linkedList) addToHead(value int){
 var node1 = node{}
 node1.value=value;	
 if lList.head!=nil {
	 node1.next=lList.head
 }
 lList.head=&node1

}


func (lList *linkedList) printList(){
	fmt.Println("Elements of list")
	for node:=lList.head;node!=nil;node=node.next{
		fmt.Println(node.value,node.next)
	}
	fmt.Println("End")
   }

   func (node *node) printNode(){
		fmt.Println(node.value,node.next)

   }

   func (lList *linkedList) lastNode() *node {
	node:=lList.head
	for ;node.next!=nil;node=node.next{		
	}
   return node
   }

   func (lList *linkedList) getNode(value int) *node {
	node1:=lList.head
	for ;node1!=nil;node1=node1.next{	
		if node1.value == value {
		return node1;	
		}
	}
   return nil
   }

   func (lList *linkedList) addToEnd(value int){
	var node1 = node{}
	node1.value=value;	
	lList.lastNode().next=&node1	
   }

   func (lList *linkedList) addAfter(value int,newValue int){
	var newNode = node{}
	newNode.value=newValue;	
	var node1 = lList.getNode(value)
	newNode.next=node1.next	
	node1.next=&newNode
   }
   

   

func exLinkedList(){
    var list1 linkedList 
	list1.addToHead(1)
    list1.printList()
  
    list1.addToHead(2)
	list1.printList()

	list1.lastNode().printNode()

	list1.addToEnd(3)
	list1.printList()

	list1.lastNode().printNode()

	list1.getNode(1).printNode()
	list1.getNode(2).printNode()
	list1.getNode(3).printNode()


	list1.printList()

	list1.addAfter(1,4)


	list1.printList()



/*output

nikhil@nikhil:~/golang$ go run src/main.go src/linked_list.go 
Hello World!
Elements of list
1 <nil>
End
Elements of list
2 &{1 <nil>}
1 <nil>
End
1 <nil>
Elements of list
2 &{1 0xc000010270}
1 &{3 <nil>}
3 <nil>
End
3 <nil>
1 &{3 <nil>}
2 &{1 0xc000010270}
3 <nil>
Elements of list
2 &{1 0xc000010270}
1 &{3 <nil>}
3 <nil>
End
Elements of list
2 &{1 0xc0000102e0}
1 &{4 0xc000010270}
4 &{3 <nil>}
3 <nil>
End

/*
}