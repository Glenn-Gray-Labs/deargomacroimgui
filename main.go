package main

import (
	"bytes"
	"github.com/Glenn-Gray-Labs/giu"
	"github.com/Glenn-Gray-Labs/giu/imgui"
	"github.com/cosmos72/gomacro/fast"
	"github.com/radovskyb/watcher"
	"io/ioutil"
	"regexp"
	"time"
)

/**********************************************************************************************************************/
// "Main" Equivalent: Execute anything you want here, be sure to return your desired layout for the master window.
/**********************************************************************************************************************/
func layout() giu.Layout {
	return giu.Layout{
		giu.Label("Edit main.go while running, and\nwatch your changes happening live."),
		giu.Line(
			giu.Button("Button One!", btnOne),
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

func reload() {
	// https://github.com/ocornut/imgui/issues/707#issuecomment-512669512
	style := imgui.CurrentStyle()
	style.SetColor(imgui.StyleColorText, imgui.Vec4{1.00, 1.00, 1.00, 1.00})
	style.SetColor(imgui.StyleColorTextDisabled, imgui.Vec4{0.36, 0.42, 0.47, 1.00})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{0.11, 0.15, 0.17, 1.00})
	style.SetColor(imgui.StyleColorChildBg, imgui.Vec4{0.15, 0.18, 0.22, 1.00})
	style.SetColor(imgui.StyleColorPopupBg, imgui.Vec4{0.08, 0.08, 0.08, 0.94})
	style.SetColor(imgui.StyleColorBorder, imgui.Vec4{0.08, 0.10, 0.12, 1.00})
	style.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{0.00, 0.00, 0.00, 0.00})
	style.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{0.20, 0.25, 0.29, 1.00})
	style.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{0.12, 0.20, 0.28, 1.00})
	style.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{0.09, 0.12, 0.14, 1.00})
	style.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{0.09, 0.12, 0.14, 0.65})
	style.SetColor(imgui.StyleColorTitleBgActive, imgui.Vec4{0.08, 0.10, 0.12, 1.00})
	style.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{0.00, 0.00, 0.00, 0.51})
	style.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{0.15, 0.18, 0.22, 1.00})
	style.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{0.02, 0.02, 0.02, 0.39})
	style.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{0.20, 0.25, 0.29, 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{0.18, 0.22, 0.25, 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{0.09, 0.21, 0.31, 1.00})
	style.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{0.28, 0.56, 1.00, 1.00})
	style.SetColor(imgui.StyleColorSliderGrab, imgui.Vec4{0.28, 0.56, 1.00, 1.00})
	style.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{0.37, 0.61, 1.00, 1.00})
	style.SetColor(imgui.StyleColorButton, imgui.Vec4{0.20, 0.25, 0.29, 1.00})
	style.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{0.28, 0.56, 1.00, 1.00})
	style.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{0.06, 0.53, 0.98, 1.00})
	style.SetColor(imgui.StyleColorHeader, imgui.Vec4{0.20, 0.25, 0.29, 0.55})
	style.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{0.26, 0.59, 0.98, 0.80})
	style.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{0.26, 0.59, 0.98, 1.00})
	style.SetColor(imgui.StyleColorSeparator, imgui.Vec4{0.20, 0.25, 0.29, 1.00})
	style.SetColor(imgui.StyleColorSeparatorHovered, imgui.Vec4{0.10, 0.40, 0.75, 0.78})
	style.SetColor(imgui.StyleColorSeparatorActive, imgui.Vec4{0.10, 0.40, 0.75, 1.00})
	style.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{0.26, 0.59, 0.98, 0.25})
	style.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{0.26, 0.59, 0.98, 0.67})
	style.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{0.26, 0.59, 0.98, 0.95})
	style.SetColor(imgui.StyleColorTab, imgui.Vec4{0.11, 0.15, 0.17, 1.00})
	style.SetColor(imgui.StyleColorTabHovered, imgui.Vec4{0.26, 0.59, 0.98, 0.80})
	style.SetColor(imgui.StyleColorTabActive, imgui.Vec4{0.20, 0.25, 0.29, 1.00})
	style.SetColor(imgui.StyleColorTabUnfocused, imgui.Vec4{0.11, 0.15, 0.17, 1.00})
	style.SetColor(imgui.StyleColorTabUnfocusedActive, imgui.Vec4{0.11, 0.15, 0.17, 1.00})
	style.SetColor(imgui.StyleColorPlotLines, imgui.Vec4{0.61, 0.61, 0.61, 1.00})
	style.SetColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{1.00, 0.43, 0.35, 1.00})
	style.SetColor(imgui.StyleColorPlotHistogram, imgui.Vec4{0.90, 0.70, 0.00, 1.00})
	style.SetColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{1.00, 0.60, 0.00, 1.00})
	style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{0.26, 0.59, 0.98, 0.35})
	style.SetColor(imgui.StyleColorDragDropTarget, imgui.Vec4{1.00, 1.00, 0.00, 0.90})
	style.SetColor(imgui.StyleColorNavHighlight, imgui.Vec4{0.26, 0.59, 0.98, 1.00})
	style.SetColor(imgui.StyleColorNavWindowingHighlight, imgui.Vec4{1.00, 1.00, 1.00, 0.70})
	style.SetColor(imgui.StyleColorNavWindowingDarkening, imgui.Vec4{0.80, 0.80, 0.80, 0.20})
	style.SetColor(imgui.StyleColorModalWindowDarkening, imgui.Vec4{0.80, 0.80, 0.80, 0.35})
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

	// The following strips the package line, and appends calls to `reload()` then `layout()` as the final interpreted
	// operations, which returns our layout as an `interface{}` that we simply cast back to a `gui.Layout` in order to
	// dog-food changes.
	layout, _ := fast.New().Eval1(string(lastRaw[len("package main\n\n"):]) + "\nreload()\nlayout()")
	lastLayout = layout.Interface().(giu.Layout)
	giu.SingleWindow("deargomacroimgui", lastLayout)
}

func main() {
	// Watch go files.
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write, watcher.Create)
	w.AddFilterHook(watcher.RegexFilterHook(regexp.MustCompile("(?i).go$"), false))
	if err := w.AddRecursive("."); err != nil {
		println(err.Error())
		w.Close()
	} else {
		// Fire up the watcher!
		go func() {
			for {
				select {
				case <-w.Event:
					giu.Update()
				case err := <-w.Error:
					println(err.Error())
					return
				case <-w.Closed:
					return
				}
			}
		}()
		go func() {
			if err := w.Start(time.Millisecond * 200); err != nil {
				println(err.Error())
				w.Close()
			}
		}()
	}

	// TODO: Read a Config File to Setup Master Window
	giu.NewMasterWindow("deargomacroimgui", 320, 200, 0, nil).Main(loop)
}
