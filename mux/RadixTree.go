package httpmux

// handlers are functions acting at certain urls
type handler func(data []byte) bool

type RadixNode struct {
	nodeMap NodeMap
	Data    string
}

type NodeMap map[string]*RadixNode

func NewRadixNode(d string) (rn *RadixNode) {
	rn.nodeMap = make(NodeMap)
	rn.Data = d
	return rn
}

/*
 * Radix Tree representing URLs
 */

type RadixTree struct {
	Node *RadixNode
}

func AddNode(key string, value NodeMap) bool {
	return true
}

func NewRadixTree(start string) (rt *RadixTree) {
	rt.Node = NewRadixNode(start, nil, nil)
	return rt
}
