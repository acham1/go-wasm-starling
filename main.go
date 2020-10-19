package main

import (
	"fmt"
	"syscall/js"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/htmlcanvas"
)

var fontFamily *canvas.FontFamily

func main() {
	cvs := js.Global().Get("document").Call("getElementById", "canvas")
	c := htmlcanvas.New(cvs, 200, 100, 5.0)

	ctx := canvas.NewContext(c)
	draw(ctx)

	alive := make(chan bool)
	<-alive
}

func draw(c *canvas.Context) {
	// Draw an elliptic arc being dashed
	ellipse, err := canvas.ParseSVG(fmt.Sprintf("A10 30 30 1 0 30 0z"))
	if err != nil {
		panic(err)
	}
	c.SetFillColor(canvas.Whitesmoke)
	c.DrawPath(110, 40, ellipse)

	c.SetFillColor(canvas.Transparent)
	c.SetStrokeColor(canvas.Black)
	c.SetStrokeWidth(0.75)
	c.SetStrokeCapper(canvas.RoundCap)
	c.SetStrokeJoiner(canvas.RoundJoin)
	c.SetDashes(0.0, 2.0, 4.0, 2.0, 2.0, 4.0, 2.0)
	//ellipse = ellipse.Dash(0.0, 2.0, 4.0, 2.0).Stroke(0.5, canvas.RoundCap, canvas.RoundJoin)
	c.DrawPath(110, 40, ellipse)
	c.SetStrokeColor(canvas.Transparent)
	c.SetDashes(0.0)

	// Draw an closed set of points being smoothed
	polyline := &canvas.Polyline{}
	polyline.Add(0.0, 0.0)
	polyline.Add(30.0, 0.0)
	polyline.Add(30.0, 15.0)
	polyline.Add(0.0, 30.0)
	polyline.Add(0.0, 0.0)
	c.SetFillColor(canvas.Seagreen)
	c.FillColor.R = byte(float64(c.FillColor.R) * 0.25)
	c.FillColor.G = byte(float64(c.FillColor.G) * 0.25)
	c.FillColor.B = byte(float64(c.FillColor.B) * 0.25)
	c.FillColor.A = byte(float64(c.FillColor.A) * 0.25)
	c.SetStrokeColor(canvas.Seagreen)
	c.DrawPath(155, 35, polyline.Smoothen())

	c.SetFillColor(canvas.Transparent)
	c.SetStrokeColor(canvas.Black)
	c.SetStrokeWidth(0.5)
	c.DrawPath(155, 35, polyline.ToPath())
	c.SetStrokeWidth(0.75)
	for _, coord := range polyline.Coords() {
		c.DrawPath(155, 35, canvas.Circle(2.0).Translate(coord.X, coord.Y))
	}

	// Draw a open set of points being smoothed
	polyline = &canvas.Polyline{}
	polyline.Add(0.0, 0.0)
	polyline.Add(20.0, 10.0)
	polyline.Add(40.0, 30.0)
	polyline.Add(60.0, 40.0)
	polyline.Add(80.0, 20.0)
	c.SetStrokeColor(canvas.Dodgerblue)
	c.DrawPath(10, 15, polyline.Smoothen())
	c.SetStrokeColor(canvas.Black)
	for _, coord := range polyline.Coords() {
		c.DrawPath(10, 15, canvas.Circle(2.0).Translate(coord.X, coord.Y))
	}
}
