package algo4go

import "testing"

func testMaxDepth(t *testing.T, F BinaryNode) {
	if MaxDepth(F) != 4 {
		t.Error("maxdepth check fail!")
	}
}

func testTraversal(t *testing.T, F BinaryNode) {
	PreOrderResult := []interface{}{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	if !testEqual(PreOrderRecursion(F, []interface{}{}), PreOrderResult) {
		t.Error("pre order recursion traversal fail!")
	}
	if !testEqual(PreOrder(F), PreOrderResult) {
		log(PreOrder(F))
		t.Error("pre order traversal fail!")
	}
	InOrderResult := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	if !testEqual(InOrderRecursion(F, []interface{}{}), InOrderResult) {
		t.Error("in order recursion traversal fail!")
	}
	if !testEqual(InOrder(F), InOrderResult) {
		t.Error("pre order traversal fail!")
	}
	PostOrderResult := []interface{}{"A", "C", "E", "D", "B", "H", "I", "G", "F"}
	if !testEqual(PostOrderRecursion(F, []interface{}{}), PostOrderResult) {
		t.Error("post order recursion traversal fail!")
	}
	if !testEqual(PostOrder(F), PostOrderResult) {
		t.Error("post order traversal fail!")
	}

}

func TestBinaryTree(t *testing.T) {
	// test case from wikipedia
	// https://en.wikipedia.org/wiki/Tree_traversal
	A := NewBinaryNode("A")
	B := NewBinaryNode("B")
	C := NewBinaryNode("C")
	D := NewBinaryNode("D")
	E := NewBinaryNode("E")
	F := NewBinaryNode("F")
	G := NewBinaryNode("G")
	H := NewBinaryNode("H")
	I := NewBinaryNode("I")
	F.left = &B
	F.right = &G
	B.left = &A
	B.right = &D
	D.left = &C
	D.right = &E
	G.right = &I
	I.left = &H
	t.Run("test max depth", func(t *testing.T) { testMaxDepth(t, F) })
	t.Run("test traversal", func(t *testing.T) { testTraversal(t, F) })
}
