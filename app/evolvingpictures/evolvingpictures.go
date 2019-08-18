// http://gameswithgo.org/balloons/balloons.zip

package main

// Experiment! draw some crazy stuff!
// Gist it next week and I'll show it off on stream
// ep16 58:20

import (
	"fmt"
	"time"
	. "github.com/lashleykeith/golanggames/app/evolvingpictures/apt"
	//. "github.com/jackmott/evolvingpictures/apt"
	"github.com/veandco/go-sdl2/sdl")

const winWidth, winHeight, winDepth int = 1280, 720, 100

type audioState struct {
	explosionBytes []byte
	deviceID       sdl.AudioDeviceID
	audioSpec      *sdl.AudioSpec
}

type mouseState struct {
	leftButton  bool
	rightButton bool
	x, y        int
}

func getMouseState() mouseState {
	mouseX, mouseY, mouseButtonState := sdl.GetMouseState()
	leftButton := mouseButtonState & sdl.ButtonLMask()
	rightButton := mouseButtonState & sdl.ButtonRMask()
	var result mouseState
	result.x = int(mouseX)
	result.y = int(mouseY)
	result.leftButton = !(leftButton == 0)
	result.rightButton = !(rightButton == 0)
	return result
}



type rgba struct {
	r, g, b byte
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func setPixel(x, y int, c rgba, pixels []byte) {
	index := (y*winWidth + x) * 4
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}

}

func pixelsToTexture(renderer *sdl.Renderer, pixels []byte, w, h int) *sdl.Texture {
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(w), int32(h))
	if err != nil {
		panic(err)
	}
	tex.Update(nil, pixels, w*4)
	return tex
}

func aptToTexture(redNode, greenNode, blueNode Node, w, h int, renderer *sdl.Renderer) *sdl.Texture {
	// -1.0 and 1.0
	scale := float32(255 / 2)
	offset := float32(-1.0 * scale)
	pixels := make([]byte, w*h*4)
	pixelIndex := 0
	for yi := 0; yi < h; yi++ {
		y := float32(yi)/float32(h)*2 - 1
		for xi := 0; xi < w; xi++ {
			x := float32(xi)/float32(w)*2 - 1

			r := redNode.Eval(x, y)
			g := greenNode.Eval(x, y)
			b := blueNode.Eval(x, y)


			pixels[pixelIndex] = byte(r*scale - offset)
			pixelIndex++
			pixels[pixelIndex] = byte(g*scale - offset)
			pixelIndex++
			pixels[pixelIndex] = byte(b*scale - offset)
			pixelIndex++
			pixelIndex++

		}
	}
	return pixelsToTexture(renderer, pixels, w, h)
}

func main() {
	sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Evolving Pictures", 200, 200,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	// Note, run "go get -u github.com/veandco/go-sdl2/sdl" to get the latest sdl2 bindings
	// which fixes an unecessary parameter in the LoadWAV call
	/*explosionBytes, audioSpec := sdl.LoadWAV("explode.wav")
	audioID, err := sdl.OpenAudioDevice("", false, audioSpec, nil, 0)
	if err != nil {
		panic(err)
	}
	defer sdl.FreeWAV(explosionBytes)

	audioState := audioState{explosionBytes, audioID, audioSpec}
	*/
	
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	
	var elapsedTime float32
	currentMouseState := getMouseState()
	// prevMouseState := currentMouseState

	aptR := GetRandomNode()
	aptG := GetRandomNode()
	aptB := GetRandomNode()

	num := rand.Intn(20)
	for i := 0; i < num; i++ {
		aptR.AddRandom(GetRandomNode())
	}

	num := rand.Intn(20)
	for i := 0; i < num; i++ {
		aptG.AddRandom(GetRandomNode())
	}

	num := rand.Intn(20)
	for i := 0; i < num; i++ {
		aptB.AddRandom(GetRandomNode())
	}

	for {
		_, nilCount := aptR.NodeCounts()
		if nilCount == 0 {
			break
		}
		aptR.AddRandom(GetRandomLeaf())
	}



	for {
		_, nilCount := aptG.NodeCounts()
		if nilCount == 0 {
			break
		}
		aptG.AddRandom(GetRandomLeaf())
	}

	for {
		_, nilCount := aptB.NodeCounts()
		if nilCount == 0 {
			break
		}
		aptB.AddRandom(GetRandomLeaf())
	}
	
	tex := aptToTexture(aptR, aptG, aptB, 640, 480, renderer)

	x := &OpX{}
	y := &OpY{}
	sine := &OpSin{}
	noise := &OpNoise{}
	atan2 := &OpMult{}
	plus := &OpPlus{}
	atan2.LeftChild = x
	atan2.RightChild = noise
	noise.LeftChild = x
	noise.RightChild = y
	sine.Child = atan2
	plus.LeftChild = y
	plus.RightChild = sine



	for {
		frameStart := time.Now()

		currentMouseState = getMouseState()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.TouchFingerEvent:
				if e.Type == sdl.FINGERDOWN {
					touchX := int(e.X * float32(winWidth))
					touchY := int(e.Y * float32(winHeight))
					currentMouseState.x = touchX
					currentMouseState.y = touchY
					currentMouseState.leftButton = true
				}
			}
		}

		renderer.Copy(tex, nil, nil)

		
		renderer.Present()
		elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)
		//	fmt.Println("ms per frame:", elapsedTime)
		if elapsedTime < 5 {
			sdl.Delay(5 - uint32(elapsedTime))
			elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)
		}

		// prevMouseState = currentMouseState
	}

}

//https://wiki.libsdl.org/SDL_GetMouseState

//https://wiki.libsdl.org/SDL_TouchFingerEvent


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

// 66
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
	return "( Sin " + op.Child.String() + ")"
}

type OpPlus struct {
	DoubleNode
}

func (op *OpPlus) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) + op.RightChild.Eval(x, y)
}

func (op *OpPlus) String() string {
	return "( + " + op.LeftChild.String() + " " + op.RightChild.String() + ")"
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
