package tree

import (
	"strconv"
	"strings"
)

// Tree is a binary tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

const terminationSymbol = "#"

// NewFromPreOrderSeq creates a new binary tree from a pre-ordered sequence of values. The terminations
// are marked with character
func NewFromPreOrderedSeq(values []string) *Tree {
	if len(values) == 0 {
		return nil
	}
	root, _ := buildPreOrderedTree(values)
	return root
}

func buildNode(value string) *Tree {
	if strings.Compare(value, terminationSymbol) == 0 {
		return nil
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}

	return &Tree{Value: v}
}

func buildPreOrderedTree(values []string) (*Tree, []string) {
	root := buildNode(values[0])
	if root == nil {
		return nil, values[1:]
	}
	leftNode, rightValues := buildPreOrderedTree(values[1:])
	root.Left = leftNode
	rightNode, remainingValues := buildPreOrderedTree(rightValues)
	root.Right = rightNode
	return root, remainingValues
}

func walkPreOrder(t *Tree, ch chan string, quit chan int) {
	if t == nil {
		ch <- terminationSymbol
		return
	}

	select {
	case ch <- strconv.Itoa(t.Value):
	case <-quit:
		return
	}
	walkPreOrder(t.Left, ch, quit)
	walkPreOrder(t.Right, ch, quit)
}

// WalkPreOrder traverses the binary tree in pre-order and sends the values in the provided channel
func WalkPreOrder(t *Tree, ch chan string, quit chan int) {
	walkPreOrder(t, ch, quit)
	close(ch)
}

func walkInOrder(t *Tree, ch chan string, quit chan int) {
	walkInOrder(t.Left, ch, quit)
	if t == nil {
		ch <- terminationSymbol
		return
	}

	select {
	case ch <- strconv.Itoa(t.Value):
	case <-quit:
		return
	}
	walkInOrder(t.Right, ch, quit)
}

// WalkInOrder traverses the binary tree in in-order and sends the values in the provided channel
func WalkInOrder(t *Tree, ch chan string, quit chan int) {
	walkInOrder(t, ch, quit)
	close(ch)
}

func walkPostOrder(t *Tree, ch chan string, quit chan int) {
	walkPostOrder(t.Left, ch, quit)
	walkPostOrder(t.Right, ch, quit)
	if t == nil {
		ch <- terminationSymbol
		return
	}

	select {
	case ch <- strconv.Itoa(t.Value):
	case <-quit:
		return
	}
}

// WalkPostOrder traverses the binary tree in post-order and sends the values in the provided channel
func WalkPostOrder(t *Tree, ch chan string, quit chan int) {
	walkPostOrder(t, ch, quit)
	close(ch)
}
