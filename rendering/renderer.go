package rendering

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer struct {
	ClearColor mgl32.Vec4
}

func NewRenderer() *Renderer {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	return &Renderer{ClearColor: mgl32.Vec4{0.0, 0.0, 0.0, 1.0}}
}

func (r *Renderer) SetClearColor(_r, g, b, a float32) {
	r.ClearColor = mgl32.Vec4{_r, g, b, a}
}

func (r *Renderer) Clear() {
	gl.ClearColor(r.ClearColor.X(), r.ClearColor.Y(), r.ClearColor.Z(), r.ClearColor.W())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

//func (r *Renderer) RenderEntities(pobjects []*phattObjects.Entity, win *Window) {
//
//}
