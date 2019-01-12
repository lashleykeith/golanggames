package apt

import (
	"fmt"
	//. "github.com/jackmott/gameswithgo-public/evolvingpictures/apt"
)

type Node interface {
	Eval(x, y float32) float32
}

type LeafNode struct{}

type SingleNode struct {
	Child Node
}

type DoubleNode struct {
	LeftChild  Node
	RightChild Node
}

type OpPlus struct {
	DoubleNode
}

func (op *OpPlus) Eval(x, y float32) float32 {
	fmt.Println(x, y)
	fmt.Println(op.LeftChild.Eval(x, y))
	fmt.Println(op.RightChild.Eval(x, y))
	return op.LeftChild.Eval(x, y) + op.RightChild.Eval(x, y)
}

type OpX LeafNode

func (op *OpX) Eval(x, y float32) float32 {
	return x
}

type OpY LeafNode

func (op *OpY) Eval(x, y float32) float32 {
	return y
}

// EP 33:00
// https://github.com/veandco/go-sdl2/
