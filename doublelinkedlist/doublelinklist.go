package doublelinkedlist

// Node Declare a simple linked list
type Node struct {
	Data int
	Next *Node
	Prev *Node
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
	item.Prev = nil
	return item
}

//InitList Initialize the list
func InitList(data int) {
	var item *Node = new(Node)
	item.Data = data
	item.Next = nil
	item.Prev = nil
	Head = item
	Trav = item
	Tail = item

}

//ResetTrav reset Trav pointer to head

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
	node.Prev = Trav
	Tail = node

}

//AddNodeToHead add item to linked list HEAD
func AddNodeToHead(data int) {
	var node *Node = CreateNode(data)
	Head.Prev = node
	node.Next = Head

	node.Prev = nil
	Head = node
}

func AddToCurrentTravPosition(data int) {
	var node *Node = CreateNode(data)
	node.Prev = Trav
	node.Next = Trav.Next
	Trav.Next.Prev = node
	Trav.Next = node

}

func DeleteNode(node *Node) {

	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	Trav = node.Prev
	node = nil
}
func ResetTrav() {
	Trav = Head
}
func GotoTail() {
	Trav = Tail
}
