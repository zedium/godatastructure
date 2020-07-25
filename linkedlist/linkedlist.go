package linkedlist

// Node Declare a simple linked list
type Node struct {
	Data int
	Next *Node
}

//Head is linkedlist head
//Tail is linkedlist tail
//Trav linkedlist trav
var Head, Tail, Trav *Node

//CreateNode Initialize the list
func CreateNode(data int) *Node {
	var item *Node = new(Node)
	item.Data = data
	item.Next = nil
	return item
}

//InitList Initialize the list
func InitList(data int) {
	var item *Node = new(Node)
	item.Data = data
	item.Next = nil
	Head = item
	Trav = item
	Tail = item

}

//ResetTrav reset Trav pointer to head
func ResetTrav() {
	Trav = Head
}

//AddNodeToTail add item to linked list TAIL
func AddNodeToTail(data int) {
	var node *Node = CreateNode(data)
	node.Next = nil

	Trav = Head
	for Trav.Next != nil {
		//fmt.Println(Trav.Data)
		Trav = Trav.Next
	}

	Trav.Next = node
	Tail = node

}

//AddNodeToHead add item to linked list HEAD
func AddNodeToHead(data int) {
	var node *Node = CreateNode(data)
	node.Next = Head
	Head = node
}
