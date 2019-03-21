package algo4go

import (
	"fmt"
	"strings"
)

// BinaryNode construct binary tree
type BinaryNode struct {
	value interface{}
	left  *BinaryNode
	right *BinaryNode
}

func (node BinaryNode) String() string {
	result := fmt.Sprintf("%v", node.value)
	if node.left != nil {
		result = fmt.Sprintf("%v-", node.left.value) + result
	} else {
		result = "nil-" + result
	}
	if node.right != nil {
		result += fmt.Sprintf("-%v", node.right.value)
	} else {
		result += "-nil"
	}
	return result
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
	return map[bool]int{true: MaxDepth(*root.left), false: MaxDepth(*root.right)}[MaxDepth(*root.left) > MaxDepth(*root.right)] + 1
}

// InOrderMorris is a binary tree travelsal algorithms with O(1) space complexity
// 1.If left child is null, Move to right child.
// .Else, find the mostright node of left subtree. Two cases arise:
// 		a) The right child of the mostright already points to the current node, set it to nil. Move to right child of current node.
// 		b) The right child of the mostright is nil. Set it to current node. Move to left child of current node.
// 2.Iterate until current node is not nil.
func InOrderMorris(node BinaryNode) []interface{} {
	results := []interface{}{}
	root := &node
	current := root
	mostRight := root
	for current != nil {
		mostRight = current.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != current {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = current
				current = current.left
			} else {
				mostRight.right = nil
				results = append(results, current.value)
				current = current.right
			}
		} else {
			results = append(results, current.value)
			current = current.right
		}
	}

	return results
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
// NOTE:In a binary search tree, in-order traversal retrieves data in sorted orde
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

// PostOrderRecursion : right left mid
func PostOrderRecursion(root BinaryNode, results []interface{}) []interface{} {
	if root.left != nil {
		results = PostOrderRecursion(*root.left, results)
	}
	if root.right != nil {
		results = PostOrderRecursion(*root.right, results)
	}
	results = append(results, root.value)
	return results
}

// PostOrder run changed preorder travel and put it's result into stack
func PostOrder(root BinaryNode) []interface{} {
	results := []interface{}{}
	stack1 := NewStack()
	stack2 := NewStack()
	current := &root
	stack1.Push(current)
	for !stack1.Empty() {
		current = stack1.Pop().(*BinaryNode)
		stack2.Push(current.value)
		if current.left != nil {
			stack1.Push(current.left)
		}
		if current.right != nil {
			stack1.Push(current.right)
		}

	}
	// mid left right to right left mid
	for !stack2.Empty() {
		results = append(results, stack2.Pop())
	}
	return results
}

// LevelTraversal traversal tree from top to bottom
func LevelTraversal(root BinaryNode) []interface{} {
	results := []interface{}{}
	// TODO:implement with queue
	// https://godoc.org/github.com/golang-collections/go-datastructures/queue
	queue1 := make([]BinaryNode, 0)
	queue2 := make([]BinaryNode, 0)
	queue1 = append(queue1, root)
	for len(queue1) > 0 {
		for len(queue1) > 0 {
			results = append(results, queue1[0].value)
			if queue1[0].left != nil {
				queue2 = append(queue2, *queue1[0].left)
			}
			if queue1[0].right != nil {
				queue2 = append(queue2, *queue1[0].right)
			}
			queue1 = queue1[1:]
		}
		queue1, queue2 = queue2, queue1
	}
	return results
}

// SerializeBinaryTree serialize binary tree to string
// Require Binary Node value type is string
func SerializeBinaryTree(node *BinaryNode) string {
	if node == nil {
		return "#"
	}
	results := make([]string, 0)
	queue1 := make([]*BinaryNode, 0)
	queue2 := make([]*BinaryNode, 0)
	queue1 = append(queue1, node)
	for len(queue1) > 0 {
		for len(queue1) > 0 {
			if queue1[0] != nil {
				results = append(results, queue1[0].value.(string))
				queue2 = append(queue2, queue1[0].left)
				queue2 = append(queue2, queue1[0].right)
			} else {
				results = append(results, "#")
			}
			queue1 = queue1[1:]
		}
		queue1, queue2 = queue2, queue1
	}
	return strings.Join(results, "_")
}

// DeserializeBinaryTree convert string to binary tree root node
func DeserializeBinaryTree(raw string) BinaryNode {
	if raw == "#" {
		return NewBinaryNode("")
	}
	queue1 := make([]*BinaryNode, 0)
	queue2 := make([]*BinaryNode, 0)
	rawSlice := strings.Split(raw, "_")
	root := NewBinaryNode(rawSlice[0])
	rawSlice = rawSlice[1:]
	queue1 = append(queue1, &root)
	for len(queue1) > 0 {
		for len(queue1) > 0 {
			c := queue1[0]
			if rawSlice[0] != "#" {
				l := NewBinaryNode(rawSlice[0])
				c.left = &l
				queue2 = append(queue2, &l)
			} else {
				c.left = nil
			}
			rawSlice = rawSlice[1:]
			if rawSlice[0] != "#" {
				r := NewBinaryNode(rawSlice[0])
				c.right = &r
				queue2 = append(queue2, &r)
			} else {
				c.right = nil
			}
			rawSlice = rawSlice[1:]
			queue1 = queue1[1:]
		}
		queue1, queue2 = queue2, queue1
	}
	return root
}
