package framework

import "syscall/js"

//Canvas holds the configuration for an HTML5 canvas
type Canvas struct {
	canvas  js.Value
	context js.Value
	height  float64
	width   float64
}

//NewCanvas takes an HTML canvas element and returns a holding struct
func NewCanvas(element js.Value) *Canvas {
	return &Canvas{
		element,
		element.Call("getContext", "2d"),
		element.Get("height").Float(),
		element.Get("width").Float(),
	}
}

//Draw draws an asset onto the canvas at a position and size measured in percentage
func (c *Canvas) Draw(a *Asset, x, y, width, height float64) {

    // if the sprite is not on the canvas it makes no sense to render it
    if x > 100 || x < (0 - width) || y > 100 || y < (0 - height) {
        return
    }

	// convert percentages to pixels
	x = x * (c.width / 100)
	y = y * (c.height / 100)
	width = width * (c.width / 100)
	height = height * (c.height / 100)

	// render the asset image onto the canvas
	c.context.Call("drawImage", a.Image, x, y, width, height)
}

//Clear clears the canvas
func (c *Canvas) Clear() {
	c.context.Call("clearRect", 0, 0, c.width, c.height)
}
