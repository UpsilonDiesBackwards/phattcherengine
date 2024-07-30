package engine

import (
	"github.com/UpsilonDiesBackwards/phattcherengine/rendering"
	"github.com/UpsilonDiesBackwards/phattcherengine/structures"
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Engine struct {
	window *rendering.Window
	scene  *structures.Scene
	shader *rendering.Shader
}

func NewPhattcherEngine(width, height int, title string, shader *rendering.Shader) (*Engine, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := rendering.NewWindow(title)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		return nil, err
	}

	scene := structures.NewScene()

	return &Engine{window: window, scene: scene, shader: shader}, nil
}

func (e *Engine) Run() {
	for !e.window.ShouldClose() {
		e.scene.Update()

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		e.scene.Render(e.shader, e.window)
		e.window.SwapBuffers()

		glfw.PollEvents()
	}
}

func (e *Engine) Terminate() {
	glfw.Terminate()
}

func (e *Engine) GetScene() *structures.Scene {
	return e.scene
}
