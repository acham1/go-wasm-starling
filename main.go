package main

import (
	"sync"
	"syscall/js"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/htmlcanvas"
)

var (
	WindowWidth  float64         = 0
	WindowHeight float64         = 0
	DPM          float64         = 10.0
	Mu           sync.Mutex      = sync.Mutex{}
	Renderer     canvas.Renderer = nil
)

func resizeBounds(this js.Value, inputs []js.Value) interface{} {
	var (
		width  = float64(inputs[0].Float())
		height = float64(inputs[1].Float())
	)
	resizeBoundsFloat64(width, height)
	return nil
}

func resizeBoundsFloat64(width, height float64) {
	WindowWidth = width
	WindowHeight = height
	val := js.Global().Get("document").Call("getElementById", "canvas")
	Renderer = htmlcanvas.New(val, WindowWidth, WindowHeight, DPM)
}

func main() {
	js.Global().Set("resizeBounds", js.FuncOf(resizeBounds))
	resizeBoundsFloat64(0, 0)

	go func() {
		for radius := 10; radius < 1000; radius++ {
			time.Sleep(time.Second)
			ctx := canvas.NewContext(Renderer)
			ctx.DrawPath(500, 250, canvas.Circle(float64(radius)))
		}
	}()

	alive := make(chan bool)
	<-alive

}
