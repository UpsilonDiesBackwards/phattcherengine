package phattObjects

import (
	"github.com/UpsilonDiesBackwards/phattcherengine/rendering"
	"github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Quat
	Model    *rendering.Model
}

func NewEntity(model *rendering.Model, pos, sca mgl32.Vec3, rot mgl32.Quat) *Entity {
	var newRot = mgl32.Quat{0, mgl32.Vec3{1, 0, 0}}

	return &Entity{
		Position: pos,
		Scale:    sca,
		Rotation: newRot,
		Model:    model,
	}
}

func (e *Entity) GetModelMatrix() mgl32.Mat4 {
	return mgl32.Translate3D(e.Position.X(), e.Position.Y(), e.Position.Z()).
		Mul4(e.Rotation.Mat4()).
		Mul4(mgl32.Scale3D(e.Scale.X(), e.Scale.Y(), e.Scale.Z()))
}

func (e *Entity) SetPosition(x, y, z float32) {
	e.Position = mgl32.Vec3{x, y, z}
}

func (e *Entity) SetScale(x, y, z float32) {
	e.Scale = mgl32.Vec3{x, y, z}
}

func (e *Entity) SetRotation(angle float32, axis mgl32.Vec3) {
	e.Rotation = mgl32.QuatRotate(angle, axis)
}

func (e *Entity) Rotate(angle float32, axis mgl32.Vec3) {
	e.Rotation = e.Rotation.Mul(mgl32.QuatRotate(angle, axis))
}

func (e *Entity) SetModel(model *rendering.Model) {
	e.Model = model
}

func (e *Entity) Update(deltaTime float32) {
	e.Position = e.Position.Add(mgl32.Vec3{0.0, 0.0, 1.0}.Mul(deltaTime))
	e.Rotation = e.Rotation.Mul(mgl32.QuatRotate(deltaTime, mgl32.Vec3{0.0, 1.0, 0.0}))
	e.Scale = e.Scale.Add(mgl32.Vec3{0.1, 0.1, 0.1}.Mul(deltaTime))
}

func (e *Entity) Draw(shader *rendering.Shader, win *rendering.Window) {
	projection := mgl32.Perspective(mgl32.DegToRad(rendering.CameraViewport.GetFov()), win.AspectRatio(), 0.1, 2000.0)

	modelMatrix := e.GetModelMatrix()
	e.Model.Draw(shader, modelMatrix, projection, rendering.CameraViewport.GetTransform())
}
