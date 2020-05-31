package main

import (
	"io/ioutil"
)

func main() {
	if _, err := ioutil.ReadFile("imgui.go"); err == nil {
		imgui_loop(imgui_init, imgui_tick, render_tick)

		/*if val, _ := fast.New().Eval1("imgui_loop(imgui_init, imgui_tick, render_tick)"); val.String() == "never" {
			imgui_loop(imgui_init, imgui_tick, render_tick)
		}*/
	}
}
