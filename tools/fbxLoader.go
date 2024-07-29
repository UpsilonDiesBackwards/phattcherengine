package tools

import (
	"fmt"
	"github.com/oakmound/ofbx"
	"os"
)

type Model struct {
	Vertices []float32
	Normals  []float32
	UVs      []float32
	Indices  [][]int
}

func LoadFBXModel(path string) (*Model, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FBX file: %w", err)
	}
	defer file.Close()

	scene, err := ofbx.Load(file)

	model := &Model{}
	mesh := scene.Meshes[0]

	for _, v := range mesh.Geometry.Vertices {
		model.Vertices = append(model.Vertices, float32(v.X()), float32(v.Y()), float32(v.Z()))
	}

	if len(mesh.Geometry.UVs) > 0 {
		for _, uv := range mesh.Geometry.UVs[0] {
			model.UVs = append(model.UVs, float32(uv.X()), float32(uv.Y()))
		}
	}

	for _, n := range mesh.Geometry.Normals {
		model.Normals = append(model.Normals, float32(n.X()), float32(n.Y()), float32(n.Z()))
	}

	for _, idx := range mesh.Geometry.Faces {
		model.Indices = append(append(model.Indices, idx))
	}

	return model, nil
}
