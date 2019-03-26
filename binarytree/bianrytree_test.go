package binarytree

import (
	"testing"

	"github.com/MarkWh1te/algo4go/util"
)

func testSerialize(t *testing.T, F BinaryNode) {
	// test empty
	if !(SerializeBinaryTree(nil) == "#") {
		t.Error("serialize empty tree fail!")
	}
	if !(DeserializeBinaryTree("#") == NewBinaryNode("")) {
		t.Error("derialize empty tree fail!")
	}
	serializedTree := SerializeBinaryTree(&F)
	if !(serializedTree == "F_B_G_A_D_#_I_#_#_C_E_H_#_#_#_#_#_#_#") {
		t.Error("serialize tree fail!")
	}
	newRoot := DeserializeBinaryTree(serializedTree)
	PreOrderResult := []interface{}{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	if !util.TestEqual(PreOrderRecursion(newRoot, []interface{}{}), PreOrderResult) {
		util.Log(PreOrderRecursion(newRoot, []interface{}{}))
		t.Error("deserialize recursion traversal fail!")
	}

}

func testMaxDepth(t *testing.T, F BinaryNode) {
	if MaxDepth(F) != 4 {
		t.Error("maxdepth check fail!")
	}
}

func testTraversal(t *testing.T, F BinaryNode) {
	PreOrderResult := []interface{}{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	if !util.TestEqual(PreOrderRecursion(F, []interface{}{}), PreOrderResult) {
		t.Error("pre order recursion traversal fail!")
	}
	if !util.TestEqual(PreOrder(F), PreOrderResult) {
		util.Log(PreOrder(F))
		t.Error("pre order traversal fail!")
	}
	InOrderResult := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	if !util.TestEqual(InOrderMorris(F), InOrderResult) {
		util.Log(InOrderMorris(F))
		t.Error("pre order recursion morris fail!")
	}

	if !util.TestEqual(InOrderRecursion(F, []interface{}{}), InOrderResult) {
		t.Error("in order recursion traversal fail!")
	}
	if !util.TestEqual(InOrder(F), InOrderResult) {
		t.Error("pre order traversal fail!")
	}
	PostOrderResult := []interface{}{"A", "C", "E", "D", "B", "H", "I", "G", "F"}
	if !util.TestEqual(PostOrderRecursion(F, []interface{}{}), PostOrderResult) {
		t.Error("post order recursion traversal fail!")
	}
	if !util.TestEqual(PostOrder(F), PostOrderResult) {
		t.Error("post order traversal fail!")
	}
	LevelTraversalResult := []interface{}{"F", "B", "G", "A", "D", "I", "C", "E", "H"}
	if !util.TestEqual(LevelTraversal(F), LevelTraversalResult) {
		util.Log(LevelTraversal(F), "level traversal")
		t.Error("level traversal fail!")
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
	t.Run("test string tree", func(t *testing.T) {
		util.Log(F, B, G, A, D, I, C, E, H)
	})
	t.Run("test max depth", func(t *testing.T) { testMaxDepth(t, F) })
	t.Run("test traversal", func(t *testing.T) { testTraversal(t, F) })
	t.Run("test serialize", func(t *testing.T) { testSerialize(t, F) })
}
