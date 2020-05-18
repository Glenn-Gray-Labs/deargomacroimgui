package main

import (
	"github.com/AllenDang/giu"
	"github.com/cosmos72/gomacro/fast"
	"io/ioutil"
)

func btnOne() {
	println("Clicked One")
}

func btnTwo() {
	println("Clicked Two")
}

func layout() giu.Layout {
	return giu.Layout{
		giu.Label("Edit main.go while running, and\nwatch your changes happening live."),
		giu.Line(
			giu.Button("#1", btnOne),
			giu.Button("#2", btnTwo)),
	}
}

func loop() {
	if raw, err := ioutil.ReadFile("main.go"); err == nil {
		layout, _ := fast.New().Eval1(string(raw[len("package main\n\n"):]) + "\nlayout()")
		giu.SingleWindow("deargomacroimgui", layout.Interface().(giu.Layout))
	}
}

func main() {
	giu.NewMasterWindow("deargomacroimgui", 320, 200, 0, nil).Main(loop)
}
