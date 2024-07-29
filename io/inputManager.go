package io

import (
	"github.com/UpsilonDiesBackwards/phattcherengine/rendering"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl64"
)

type KeyAction int

// UserInput Types of user input
type UserInput struct {
	InitialAction bool // Keyboard

	cursor         mgl64.Vec2 // Mouse
	cursorChange   mgl64.Vec2 //
	cursorLast     mgl64.Vec2 //
	bufferedChange mgl64.Vec2 //
}

// Key Actions
const (
	NO_ACTION = iota

	// Viewport
	VP_FORW
	VP_BACK
	VP_LEFT
	VP_RGHT
	VP_UP
	VP_DOWN

	// Editor
	ED_QUIT
)

var ActionState = make(map[KeyAction]bool)

var keyToActionMap = map[glfw.Key]KeyAction{
	glfw.KeyW:     VP_FORW,
	glfw.KeyS:     VP_BACK,
	glfw.KeyA:     VP_LEFT,
	glfw.KeyD:     VP_RGHT,
	glfw.KeySpace: VP_UP,
	glfw.KeyC:     VP_DOWN,

	glfw.KeyEscape: ED_QUIT,
}

func InputManager(aW *rendering.Window, uI *UserInput) {
	aW.SetKeyCallback(KeyCallBack)
	aW.SetCursorPosCallback(uI.MouseCallBack)
}

func KeyCallBack(aW *glfw.Window, key glfw.Key, scancode int, action glfw.Action, modifier glfw.ModifierKey) {
	a, ok := keyToActionMap[key]
	if !ok {
		return
	}

	switch action {
	case glfw.Press: // Key was pressed
		ActionState[a] = true
	case glfw.Release: // Key was released
		ActionState[a] = false
	}
}

func (cInput UserInput) Cursor() mgl64.Vec2 { return cInput.cursor }

func (cInput UserInput) CursorChange() mgl64.Vec2 { return cInput.cursorChange }

func (cInput *UserInput) CheckpointCursorChange() {
	cInput.cursorChange[0] = cInput.bufferedChange[0]
	cInput.cursorChange[1] = cInput.bufferedChange[1]

	cInput.cursor[0] = cInput.cursor[0]
	cInput.cursor[1] = cInput.cursor[1]

	cInput.bufferedChange[0] = 0
	cInput.bufferedChange[1] = 0
}

func (cInput *UserInput) MouseCallBack(aW *glfw.Window, xpos, ypos float64) {
	if cInput.InitialAction {
		cInput.cursorLast[0] = xpos
		cInput.cursorLast[1] = ypos
		cInput.InitialAction = false
	}

	cInput.bufferedChange[0] += xpos - cInput.cursorLast[0]
	cInput.bufferedChange[1] += ypos - cInput.cursorLast[1]

	cInput.cursorLast[0] = xpos
	cInput.cursorLast[1] = ypos
}
