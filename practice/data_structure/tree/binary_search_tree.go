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

func (n *node) exist(data int) bool {
	if n == nil {
		return false
	}
	if n.data == data {
		return true
	}
	if n.data > data {
		return n.left.exist(data)
	} else {
		return n.right.exist(data)
	}
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

func (bst *binarySearchTree) exist(data int) bool {
	return bst.top.exist(data)
}

func (bst *binarySearchTree) traverse(callback func(int)) {
	bst.top.traverse(callback)
}
