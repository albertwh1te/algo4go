package algo4go

// BinaryNode construct binary tree
type BinaryNode struct {
	value interface{}
	left  *BinaryNode
	right *BinaryNode
}

// NewBinaryNode construct new binary node with no child
func NewBinaryNode(v interface{}) BinaryNode {
	return BinaryNode{value: v, left: nil, right: nil}
}

// MaxDepth return the max depth of a tree
func MaxDepth(root BinaryNode) int {
	if root.left == nil && root.right == nil {
		return 1
	}
	if root.left == nil {
		return MaxDepth(*root.right) + 1
	}
	if root.right == nil {
		return MaxDepth(*root.left) + 1
	}
	return MaxDepth(*root.left) + MaxDepth(*root.right) + 1
}

// PreOrderRecursion parent node is processed before any of its child nodes is done.
func PreOrderRecursion(root BinaryNode, results []interface{}) []interface{} {
	results = append(results, root.value)
	if root.left != nil {
		results = PreOrderRecursion(*root.left, results)
	}
	if root.right != nil {
		results = PreOrderRecursion(*root.right, results)
	}
	return results
}

// PreOrder push right node to stack first , and push left  node to stack
func PreOrder(root BinaryNode) []interface{} {
	results := []interface{}{}
	current := &root
	stack := NewStack()
	stack.Push(current)
	for !stack.Empty() {
		current := stack.Pop().(*BinaryNode)
		results = append(results, current.value)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
	return results
}

// InOrderRecursion is a topologically sorted one, because a parent node is processed before any of its child nodes is done.
func InOrderRecursion(root BinaryNode, results []interface{}) []interface{} {
	if root.left != nil {
		results = InOrderRecursion(*root.left, results)
	}
	results = append(results, root.value)
	if root.right != nil {
		results = InOrderRecursion(*root.right, results)
	}
	return results
}

// InOrder left mid right
func InOrder(root BinaryNode) []interface{} {
	results := []interface{}{}
	stack := NewStack()
	current := &root
	for !stack.Empty() || current != nil {
		if current != nil {
			stack.Push(current)
			current = current.left
		} else {
			current = stack.Pop().(*BinaryNode)
			results = append(results, current.value)
			current = current.right
		}
	}
	return results
}
