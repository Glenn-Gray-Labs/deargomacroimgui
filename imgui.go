// +build glfw

package main

import (
	"bytes"
	"fmt"
	"github.com/inkyblackness/imgui-go/v2"
	"io/ioutil"
	"os"
)

const (
	mouseButtonPrimary   = 0
	mouseButtonSecondary = 1
	mouseButtonTertiary  = 2
)

var showDemoWindow = true
var clearColor = [3]float32{0.0, 0.0, 0.0}
var f = float32(0)
var counter = 0
var showAnotherWindow = false

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type Renderer interface {
}

func render_tick(platform Platform, renderer Renderer) {
	//
}

func imgui_tick() {

	// 1. Show a simple window.
	// Tip: if we don't call imgui.Begin()/imgui.End() the widgets automatically appears in a window called "Debug".
	{
		imgui.Text("ภาษาไทย测试조선말")               // To display these, you'll need to register a compatible font
		imgui.Text("Hello, world!")              // Display some text
		imgui.SliderFloat("float", &f, 0.0, 1.0) // Edit 1 float using a slider from 0.0f to 1.0f

		imgui.Checkbox("Demo Window", &showDemoWindow) // Edit bools storing our window open/close state
		imgui.Checkbox("Another Window", &showAnotherWindow)

		if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
			counter++
		}
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("counter = %d", counter))
	}

	// 2. Show another simple window. In most cases you will use an explicit Begin/End pair to name your windows.
	if showAnotherWindow {
		// Pass a pointer to our bool variable (the window will have a closing button that will clear the bool when clicked)
		imgui.BeginV("Another window", &showAnotherWindow, 0)
		imgui.Text("Hello from another window!")
		if imgui.Button("Close Me") {
			showAnotherWindow = false
		}
		imgui.End()
	}

	// 3. Show the ImGui demo window. Most of the sample code is in imgui.ShowDemoWindow().
	// Read its code to learn more about Dear ImGui!
	if showDemoWindow {
		// Normally user code doesn't need/want to call this because positions are saved in .ini file anyway.
		// Here we just want to make the demo initial state a bit more friendly!
		const demoX = 650
		const demoY = 20
		imgui.SetNextWindowPosV(imgui.Vec2{X: demoX, Y: demoY}, imgui.ConditionFirstUseEver, imgui.Vec2{})

		imgui.ShowDemoWindow(&showDemoWindow)
	}
}

/**********************************************************************************************************************/
// Read-Only: This is the engine that allows you to rapidly prototype. Changes will only be picked up in a new build.
/**********************************************************************************************************************/
var lastRaw []byte

func imgui_loop(imgui_init func(), imgui_tick func(), render_tick func(platform Platform, renderer Renderer)) {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := NewGLFW(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	//imgui.CurrentIO().SetClipboard(platform)

	for !platform.ShouldStop() {
		// Try to Read and Evaluate Sources: If nothing has changed, do nothing... else, update!
		if raw, err := ioutil.ReadFile("imgui.go"); err != nil || bytes.Equal(lastRaw, raw) {
			// nothing to do.
		} else {
			lastRaw = raw
			imgui_init()
		}

		// Input Events
		platform.ProcessEvents()

		// Signal start of a new frame
		platform.NewFrame()
		imgui.NewFrame()
		imgui_tick()
		imgui.Render()

		renderer.PreRender(clearColor)
		render_tick(platform, renderer)
		renderer.Render(platform.DisplaySize(), platform.FramebufferSize(), imgui.RenderedDrawData())
		platform.PostRender()
	}
}

/*func (platform Platform) Text() (string, error) {
	return platform.ClipboardText()
}

func (platform Platform) SetText(text string) {
	platform.SetClipboardText(text)
}*/

/**********************************************************************************************************************/
// "Main" Equivalent: Execute anything you want here, be sure to return your desired layout for the master window.
/**********************************************************************************************************************/

// Example Callbacks
func btnOne() {
	println("Clicked One.")
}

func btnTwo() {
	println("Clicked Two!")
}

func imgui_init() {
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
