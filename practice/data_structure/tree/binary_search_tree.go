package tree

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) insert(data int) {
	if n.data == data {
		return
	}

	if n.data > data {
		if n.left == nil {
			n.left = &node{data: data}
			return
		}
		n.left.insert(data)
		return
	}

	if n.right == nil {
		n.right = &node{data: data}
		return
	}
	n.right.insert(data)
}

func (n *node) traverse(callback func(int)) {
	if n == nil {
		return
	}
	n.left.traverse(callback)
	callback(n.data)
	n.right.traverse(callback)
}

type binarySearchTree struct {
	top *node
}

func (bst *binarySearchTree) insert(data int) {
	if bst.top == nil {
		bst.top = &node{data: data}
		return
	}
	bst.top.insert(data)
}

//func (bst *binarySearchTree) pop() int {
//	return 0
//}

func (bst *binarySearchTree) traverse(callback func(int)) {
	bst.top.traverse(callback)
}
