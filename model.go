package list

type List struct {
	First *Node
	Last  *Node
	Size  int
}

type Node struct {
	Prev  *Node
	Next  *Node
	Value int
}
