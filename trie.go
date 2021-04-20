package dfa

type Node struct {
	IsEnd bool
	Value string
	Child map[rune]*Node
}

func newNode(val string) *Node {
	return &Node{
		IsEnd: false,
		Value: val,
		Child: make(map[rune]*Node),
	}
}

type Trie struct {
	root *Node
	size int
}

func (t *Trie) Root() *Node {
	return t.root
}

func (t *Trie) Insert(key string) {
	curNode := t.root
	for _, v := range key {
		if curNode.Child[v] == nil {
			curNode.Child[v] = newNode(string(v))
		}
		curNode = curNode.Child[v]
		curNode.Value = string(v)
	}

	if !curNode.IsEnd {
		t.size++
		curNode.IsEnd = true
	}
}

func (t *Trie) PrefixMatch(key string) []string {
	node, _ := t.findNode(key)
	if node == nil {
		return nil
	}
	return t.Walk(node)
}

func (t *Trie) Walk(node *Node) (ret []string) {
	if node.IsEnd {
		ret = append(ret, node.Value)
	}
	for _, v := range node.Child {
		ret = append(ret, t.Walk(v)...)
	}
	return
}

func (t *Trie) findNode(key string) (node *Node, index int) {
	curNode := t.root
	f := false
	for k, v := range key {
		if f {
			index = k
			f = false
		}
		if curNode.Child[v] == nil {
			return nil, index
		}
		curNode = curNode.Child[v]
		if curNode.IsEnd {
			f = true
		}
	}

	if curNode.IsEnd {
		index = len(key)
	}

	return curNode, index
}

func (t *Trie) Child(key string) *Node {
	node, _ := t.findNode(key)
	return node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(""),
		size: 0,
	}
}
