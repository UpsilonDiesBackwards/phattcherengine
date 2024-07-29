package io

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/phattcherengine/rendering"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var ViewportTransform mgl32.Mat4
var u = &UserInput{}

func InputRunner(win *rendering.Window, deltaTime float64) error {
	c := &rendering.CameraViewport
	adjCSpeed := deltaTime * float64(c.Speed)

	if ActionState[VP_FORW] {
		c.Position = c.Position.Add(c.Front.Mul(adjCSpeed))
	}

	if ActionState[VP_BACK] {
		c.Position = c.Position.Sub(c.Front.Mul(adjCSpeed))
	}
	if ActionState[VP_LEFT] {
		c.Position = c.Position.Sub(c.Front.Cross(c.Up).Mul(adjCSpeed))
	}
	if ActionState[VP_RGHT] {
		c.Position = c.Position.Add(c.Front.Cross(c.Up).Mul(adjCSpeed))
	}
	if ActionState[VP_UP] {
		c.Position = c.Position.Add(c.Up.Mul(adjCSpeed))
	}
	if ActionState[VP_DOWN] {
		c.Position = c.Position.Sub(c.Up.Mul(adjCSpeed))
	}

	if ActionState[ED_QUIT] {
		fmt.Println("Exiting!")
		glfw.Terminate()
	}

	// Cursor transform
	ViewportTransform = c.GetTransform()
	c.UpdateDirection(u)
	u.CheckpointCursorChange()

	InputManager(win, u)
	return nil
}
