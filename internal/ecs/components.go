package ecs

import (
	"github.com/cxncxl/gogame/internal/math"
)

type Component interface {
    ComponentId() ComponentId
    Clone()       Component
}

// --- Shared ------------

var emptyComponents = map[ComponentId]Component {
    BoxRendererComponentId: BoxRendererComponent{},
    TransformComponentId: TransformComponent{},
    PlayerComponentId: PlayerComponent{},
    SpriteRendererComponentId: SpriteRendererComponent{},
}

// --- Render ------------

type BoxRendererComponent struct {
    Color [3]uint8
}

var BoxRendererComponentId = Id{ raw: 3331 }
func (BoxRendererComponent) ComponentId() ComponentId {
    return BoxRendererComponentId
}
func (self BoxRendererComponent) Clone() Component {
    return BoxRendererComponent{
        Color: [3]uint8{ self.Color[0], self.Color[1], self.Color[2] },
    }
}

type SpriteRendererComponent struct {
    SpritePath string
}

var SpriteRendererComponentId = Id{ raw: 3332 }
func (SpriteRendererComponent) ComponentId() ComponentId {
    return SpriteRendererComponentId
}
func (self SpriteRendererComponent) Clone() Component {
    return SpriteRendererComponent{
        SpritePath: self.SpritePath,
    }
}

// --- Position ----------

type TransformComponent struct {
    Position *math.Vector2
    Rotation *math.Vector2
    Scale    *math.Vector2
}

var TransformComponentId = Id{ raw: 3333 }
func (TransformComponent) ComponentId() ComponentId {
    return TransformComponentId
}
func (self TransformComponent) Clone() Component {
    return TransformComponent{
        Position: &math.Vector2{
            X: self.Position.X,
            Y: self.Position.Y,
        },
        Rotation: &math.Vector2{
            X: self.Rotation.X,
            Y: self.Rotation.Y,
        },
        Scale: &math.Vector2{
            X: self.Scale.X,
            Y: self.Scale.Y,
        },
    }
}

// --- Player -----------

type PlayerComponent struct {}

var PlayerComponentId = Id{ raw: 3334 }
func (PlayerComponent) ComponentId() ComponentId {
    return PlayerComponentId
}
func (self PlayerComponent) Clone() Component {
    return PlayerComponent{}
}
