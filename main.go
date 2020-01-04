package main

import "fmt"
import "./framework"
import "./game"
import "syscall/js"

func main() {

	fmt.Println("Starting web assembly main module")
	runForever := make(chan bool)

	//    js.Global().Set("example", js.FuncOf(example))
	//    func example(this js.Value, args []js.Value) interface{} {

	element := js.Global().Get("document").Call("getElementById", "canvas")
	canvas := framework.NewCanvas(element)
	go game.Start(canvas)

	<-runForever
}
