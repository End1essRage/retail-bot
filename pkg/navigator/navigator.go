package navigator

type Tree []Node

type Node struct {
	Id     int
	Name   string
	Parent int
	Child  []int
}

func main() {
	tree := make(Tree, 0)
	tree = append(tree, Node{Id: 1, Name: "A", Child: []int{3, 5}})
	tree = append(tree, Node{Id: 2, Name: "B", Child: []int{6}})
	tree = append(tree, Node{Id: 3, Name: "C", Child: []int{4}, Parent: 1})
	tree = append(tree, Node{Id: 4, Name: "E", Parent: 3})
	tree = append(tree, Node{Id: 5, Name: "D", Parent: 1})
	tree = append(tree, Node{Id: 6, Name: "F", Child: []int{3, 5}, Parent: 2})
	tree = append(tree, Node{Id: 7, Name: "J", Parent: 6})
	tree = append(tree, Node{Id: 8, Name: "Q", Child: []int{3, 5}, Parent: 6})
	tree = append(tree, Node{Id: 9, Name: "Z", Parent: 8})
}
