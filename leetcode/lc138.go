package leetcode

var cacheNodeList map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cacheNodeList = make(map[*Node]*Node, 0)
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if val, exist := cacheNodeList[node]; exist {
		return val
	}
	newNode := &Node{
		Val: node.Val,
	}
	cacheNodeList[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
