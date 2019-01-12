package main

import (
	"fmt"

	. "github.com/lashleykeith/golanggames/app/evolvingpictures/apt"
	//. "github.com/jackmott/gameswithgo-public/evolvingpictures/apt"
)

func main() {
	x := &OpX{}
	y := &OpY{}
	plus := &OpPlus{}
	sine := &OpSin{}
	sine.Child = x
	plus.LeftChild = sine
	plus.RightChild = y

	fmt.Println(plus.Eval(5, 2))
	fmt.Println(plus)

}

//42:00

/// 63
/*
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
*/

// 64evolvingpicturesmath
/*
package apt

import (
	"math"
	//. "github.com/jackmott/gameswithgo-public/evolvingpictures/apt"
)

// + / * - Sin Cos Atan SimplexNoise X Y Constants...
// Leaf Node (0 children)
// Single Node (sin/cos)
// DoubleNode (+, -)

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

type OpSin struct {
	SingleNode
}

type OpPlus struct {
	DoubleNode
}

func (op *OpSin) Eval(x, y float32) float32 {
	return float32(math.Sin(float64(op.Child.Eval(x, y))))
}

func (op *OpPlus) Eval(x, y float32) float32 {
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
*/

// 65

/*
package apt

import (
	"math"
	//. "github.com/jackmott/gameswithgo-public/evolvingpictures/apt"
)

// + / * - Sin Cos Atan SimplexNoise X Y Constants...
// Leaf Node (0 children)
// Single Node (sin/cos)
// DoubleNode (+, -)

type Node interface {
	Eval(x, y float32) float32
	String() string
}

type LeafNode struct{}

type SingleNode struct {
	Child Node
}

type DoubleNode struct {
	LeftChild  Node
	RightChild Node
}

type OpSin struct {
	SingleNode
}

func (op *OpSin) Eval(x, y float32) float32 {
	return float32(math.Sin(float64(op.Child.Eval(x, y))))
}

func (op *OpSin) String() string {
	return "Sin(" + op.Child.String() + ")"
}

type OpPlus struct {
	DoubleNode
}

func (op *OpPlus) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) + op.RightChild.Eval(x, y)
}

func (op *OpPlus) String() string {
	return op.LeftChild.String() + " + " + op.RightChild.String()
}

type OpX LeafNode

func (op *OpX) Eval(x, y float32) float32 {
	return x
}

func (op *OpX) String() string {
	return "X"
}

type OpY LeafNode

func (op *OpY) Eval(x, y float32) float32 {
	return y
}

func (op *OpY) String() string {
	return "Y"
}

// EP 43:39
// https://github.com/veandco/go-sdl2/


*/
