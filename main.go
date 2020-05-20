package main

import (
	"bytes"
	"github.com/Glenn-Gray-Labs/giu"
	"github.com/cosmos72/gomacro/fast"
	"io/ioutil"
)

/**********************************************************************************************************************/
// "Main" Equivalent: Execute anything you want here, be sure to return your desired layout for the master window.
/**********************************************************************************************************************/
func layout() giu.Layout {
	return giu.Layout{
		giu.Label("Edit main.go while running, and\nwatch your changes happening live."),
		giu.Line(
			giu.Button("Button One.", btnOne),
			giu.Button("Button Two!", btnTwo)),
	}
}

// Example Callbacks
func btnOne() {
	println("Clicked One.")
}

func btnTwo() {
	println("Clicked Two!")
}

/**********************************************************************************************************************/
// Read-Only: This is the engine that allows you to rapidly prototype. Changes will only be picked up in a new build.
/**********************************************************************************************************************/
var lastRaw []byte
var lastLayout giu.Layout

func loop() {
	// Try to Read and Evaluate Sources: If nothing has changed, do nothing... else, update!
	if raw, err := ioutil.ReadFile("main.go"); err != nil || bytes.Equal(lastRaw, raw) {
		giu.SingleWindow("deargomacroimgui", lastLayout)
		return
	} else {
		lastRaw = raw
	}

	// The following strips the package line, and appends a call to `layout()` as the final interpreted operation, which
	// returns our layout as an `interface{}` that we simply cast back to a `gui.Layout` in order to dog-food changes.
	layout, _ := fast.New().Eval1(string(lastRaw[len("package main\n\n"):]) + "\nlayout()")
	lastLayout = layout.Interface().(giu.Layout)
	giu.SingleWindow("deargomacroimgui", lastLayout)
}

func main() {
	// TODO: Read a Config File to Setup Master Window
	giu.NewMasterWindow("deargomacroimgui", 320, 200, 0, nil).Main(loop)
}
