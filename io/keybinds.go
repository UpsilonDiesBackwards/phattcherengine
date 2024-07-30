package io

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/phattcherengine/rendering"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var ViewportTransform mgl32.Mat4
var u = &rendering.UserInput{}

func InputRunner(win *rendering.Window, deltaTime float64) error {
	c := &rendering.CameraViewport
	adjCSpeed := deltaTime * float64(c.Speed)

	if rendering.ActionState[rendering.VP_FORW] {
		c.Position = c.Position.Add(c.Front.Mul(adjCSpeed))
	}

	if rendering.ActionState[rendering.VP_BACK] {
		c.Position = c.Position.Sub(c.Front.Mul(adjCSpeed))
	}
	if rendering.ActionState[rendering.VP_LEFT] {
		c.Position = c.Position.Sub(c.Front.Cross(c.Up).Mul(adjCSpeed))
	}
	if rendering.ActionState[rendering.VP_RGHT] {
		c.Position = c.Position.Add(c.Front.Cross(c.Up).Mul(adjCSpeed))
	}
	if rendering.ActionState[rendering.VP_UP] {
		c.Position = c.Position.Add(c.Up.Mul(adjCSpeed))
	}
	if rendering.ActionState[rendering.VP_DOWN] {
		c.Position = c.Position.Sub(c.Up.Mul(adjCSpeed))
	}

	if rendering.ActionState[rendering.ED_QUIT] {
		fmt.Println("Exiting!")
		glfw.Terminate()
	}

	// Cursor transform
	ViewportTransform = c.GetTransform()
	c.UpdateDirection(u)
	u.CheckpointCursorChange()

	rendering.InputManager(win, u)
	return nil
}
