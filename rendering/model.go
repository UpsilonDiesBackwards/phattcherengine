package rendering

import (
	"fmt"
	"github.com/UpsilonDiesBackwards/phattcherengine/tools"
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type PerspectiveBlock struct {
	Project *mgl32.Mat4
	Camera  *mgl32.Mat4
	Model   *mgl32.Mat4
}

type Model struct {
	Vertices  []float32
	Indices   []uint32
	Normals   []float32
	TexCoords []float32

	vao uint32
	ubo uint32

	Texture uint32
}

func NewModel(toolModel *Model, texturePath string) (*Model, error) {
	var texture uint32
	var err error

	if texturePath == "" {
		texture = tools.CreateWhiteTexture()
	} else {
		texture, err = tools.LoadTexture(texturePath)
		if err != nil {
			return nil, err
		}
	}
	model := &Model{
		Vertices:  toolModel.Vertices,
		Indices:   toolModel.Indices,
		Normals:   toolModel.Normals,
		TexCoords: toolModel.TexCoords,

		Texture: texture,
	}
	model.setup()
	return model, nil
}

func (m *Model) setup() {
	var vbo, ebo uint32

	gl.GenVertexArrays(1, &m.vao)
	gl.BindVertexArray(m.vao)

	// Combine vertices, texCoords, and normals
	var vertexData []float32
	for i := 0; i < len(m.Vertices); i += 3 {
		vertexData = append(vertexData, m.Vertices[i], m.Vertices[i+1], m.Vertices[i+2])
		if len(m.TexCoords) > 0 {
			vertexData = append(vertexData, m.TexCoords[i%len(m.TexCoords)], m.TexCoords[(i%len(m.TexCoords))+1])
		}
		if len(m.Normals) > 0 {
			vertexData = append(vertexData, m.Normals[i%len(m.Normals)], m.Normals[(i%len(m.Normals))+1], m.Normals[(i%len(m.Normals))+2])
		}
	}

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData)*4, gl.Ptr(vertexData), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.Indices)*4, gl.Ptr(m.Indices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &m.ubo)
	gl.BindBuffer(gl.UNIFORM_BUFFER, m.ubo)
	gl.BufferData(gl.UNIFORM_BUFFER, 3*16*4, nil, gl.DYNAMIC_DRAW)
	gl.BindBufferBase(gl.UNIFORM_BUFFER, 1, m.ubo)

	stride := int32(8 * 4)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, stride, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, stride, gl.PtrOffset(5*4))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

func (m *Model) Draw(shader *Shader, modelMatrix mgl32.Mat4, project, camera mgl32.Mat4) {
	gl.BindVertexArray(m.vao)

	gl.BindBuffer(gl.UNIFORM_BUFFER, m.ubo)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 0, 16*4, gl.Ptr(&project[0]))
	gl.BufferSubData(gl.UNIFORM_BUFFER, 16*4, 16*4, gl.Ptr(&camera[0]))
	gl.BufferSubData(gl.UNIFORM_BUFFER, 32*4, 16*4, gl.Ptr(&modelMatrix[0]))
	gl.BindBufferBase(gl.UNIFORM_BUFFER, 1, m.ubo)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, m.Texture)
	shader.SetInt("texture1", 0)

	shader.Use()

	gl.DrawElements(gl.TRIANGLES, int32(len(m.Indices)), gl.UNSIGNED_INT, gl.PtrOffset(0))

	fmt.Println("coords: ", m.TexCoords)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}
