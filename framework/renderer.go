package framework

import "sync"

//newSpriteList sets the maximum number of sprites which can be drawn in one frame
func newSpriteList() []*Sprite {
	return make([]*Sprite, 0, 8192)
}

//Renderer wraps the canvas and stacks up sprites to be drawn onto the canvas
type Renderer struct {
	canvas              *Canvas
	spritesToBeRendered []*Sprite
    mux *sync.Mutex
}

//NewRenderer returns a renderer that wraps a canvas
func NewRenderer(canvas *Canvas) *Renderer {
	return &Renderer{
		canvas,
		newSpriteList(),
        &sync.Mutex{},
	}
}

//Adds a sprite to the top of the stack to be rendered
func (r *Renderer) AddSprite(s *Sprite) int {

	r.spritesToBeRendered = append(r.spritesToBeRendered, s)
	return len(r.spritesToBeRendered)
}

//Render draws all of the sprites' assets onto the canvas
func (r *Renderer) Render() bool {

    // we don't want anything writing to the sprites list when rendering
    r.mux.Lock()
    defer r.mux.Unlock()

	// first we clear the canvas
	r.canvas.Clear()

	// then we draw the sprites onto the canvas
	for _, sprite := range r.spritesToBeRendered {

		r.canvas.Draw(sprite.Asset, sprite.PositionX, sprite.PositionY, sprite.Width, sprite.Height)
	}

	// then we clear the sprites list ready for the next frame
	r.spritesToBeRendered = newSpriteList()

	return true
}

//RenderForever renders the renderer using the window's animation frame scheduler
func (r *Renderer) RenderForever() {
    r.Render()
    WindowAnimationFrame(r.RenderForever)
}
