package framework

import "syscall/js"

//callback is a typeable to allow us to execute an arbitary function
type AnimationCallback func(...interface{}) interface{}

//WindowAnimationFrame will set a callback that fires when the browser renders a
// new animation. Typically this happens at a highly variable rate, based on GPU load.
// Normally it clocks at between 15 and 60 times a second.
// Only one callback can be registered with this function, which is designed soley to
// render sprites onto the screen.
func WindowAnimationFrame(callback AnimationCallback) {

    jsCallback := func(this js.Value, args []js.Value) interface{} {
        callback()
        return true
    }

    js.Global().Call("requestAnimationFrame", js.FuncOf(jsCallback))
    return
}
