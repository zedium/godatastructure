package main

import (
	"fmt"

	"github.com/zedium/godatastructure/doublelinkedlist"
)

func main() {
	fmt.Println("Beginning...")

	doublelinkedlist.InitList(1)
	doublelinkedlist.AddNodeToHead(0)
	doublelinkedlist.AddNodeToTail(2)
	doublelinkedlist.AddNodeToTail(3)
	doublelinkedlist.AddNodeToTail(4)
	doublelinkedlist.AddNodeToHead(-1)
	doublelinkedlist.AddNodeToTail(5)
	doublelinkedlist.AddNodeToHead(-2)

	//Add item between third and fourth item
	doublelinkedlist.ResetTrav()
	for i := 0; i < 2; i++ {
		doublelinkedlist.Trav = doublelinkedlist.Trav.Next
	}
	doublelinkedlist.AddToCurrentTravPosition(255)

	fmt.Println("Printing list")
	doublelinkedlist.ResetTrav()
	for doublelinkedlist.Trav != nil {
		fmt.Println("Item data is: ", doublelinkedlist.Trav.Data)
		doublelinkedlist.Trav = doublelinkedlist.Trav.Next

	}
	//delete fourth position node
	doublelinkedlist.ResetTrav()
	for i := 0; i < 3; i++ {
		doublelinkedlist.Trav = doublelinkedlist.Trav.Next
	}
	doublelinkedlist.DeleteNode(doublelinkedlist.Trav)

	fmt.Println("Printing list in revers")
	doublelinkedlist.GotoTail()
	for doublelinkedlist.Trav != nil {
		fmt.Println("Item data is: ", doublelinkedlist.Trav.Data)
		doublelinkedlist.Trav = doublelinkedlist.Trav.Prev

	}

	fmt.Println("Ending...")
}
