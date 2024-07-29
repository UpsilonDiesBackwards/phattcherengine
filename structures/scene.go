package structures

import (
	"github.com/UpsilonDiesBackwards/phattcherengine/phattObjects"
	"github.com/UpsilonDiesBackwards/phattcherengine/rendering"
	"time"
)

type Scene struct {
	entities []phattObjects.Entity
	lastTime time.Time
}

func NewScene() *Scene {
	return &Scene{
		entities: []phattObjects.Entity{},
		lastTime: time.Now(),
	}
}

func (s *Scene) AddEntity(entity phattObjects.Entity) {
	s.entities = append(s.entities, entity)
}

func (s *Scene) Update() {
	currentTime := time.Now()
	deltaTime := float32(currentTime.Sub(s.lastTime).Seconds())
	s.lastTime = currentTime

	for i := range s.entities {
		s.entities[i].Update(deltaTime)
	}
}

func (s *Scene) Render(shader *rendering.Shader, win *rendering.Window) {
	for i := range s.entities {
		s.entities[i].Draw(shader, win)
	}
}
