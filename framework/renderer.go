package framework

import "sync"

//Renderer wraps the canvas and stacks up worlds to be drawn onto the canvas
type Renderer struct {
	canvas				*Canvas
	worldsToBeRendered	[]*World
	mux                 *sync.Mutex
}

//NewRenderer returns a renderer that wraps a canvas
func NewRenderer(canvas *Canvas) *Renderer {
	return &Renderer{
		canvas,
		make([]*World, 0, 8192),
		&sync.Mutex{},
	}
}

//Adds a sprite to the top of the stack to be rendered
func (r *Renderer) AddWorld(w *World) int {

	r.worldsToBeRendered = append(r.worldsToBeRendered, w)
	return len(r.worldsToBeRendered)
}

//Render draws all of the sprites' assets onto the canvas
func (r *Renderer) Render() bool {

	// we don't want anything writing to the sprites list when rendering
	r.mux.Lock()
	defer r.mux.Unlock()

	// first we clear the canvas
	r.canvas.Clear()

	// then for each world we draw all the sprites onto the canvas
	for _, world := range r.worldsToBeRendered {

		for _, sprite := range world.SpriteList {

			r.canvas.Draw(sprite.Asset, sprite.PositionX, sprite.PositionY, sprite.Width, sprite.Height)
		}
	}

	//	// then we clear the sprites list ready for the next frame
	//	r.spritesToBeRendered = newSpriteList()

	return true
}

//RenderForever renders the renderer using the window's animation frame scheduler
func (r *Renderer) RenderForever(...interface{}) interface{} {
	r.Render()
	defer WindowAnimationFrame(r.RenderForever)
	return true
}
