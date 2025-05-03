package ecs

import "github.com/cxncxl/gogame/internal/math"

// --- Shared ------------

var emptyComponents = map[ComponentId]Component {
    BoxRendererComponentId: BoxRendererComponent{},
    TransformComponentId: TransformComponent{},
    PlayerComponentId: PlayerComponent{},
}

// --- Render ------------

type BoxRendererComponent struct {
    Color [3]uint8
}

var BoxRendererComponentId = GetWorld().GetNewId()
func (BoxRendererComponent) ComponentId() ComponentId {
    return BoxRendererComponentId
}

type SpriteRendererComponent struct {
    SpritePath string
}

var SpriteRendererComponentId = GetWorld().GetNewId()
func (SpriteRendererComponent) ComponentId() ComponentId {
    return SpriteRendererComponentId
}

// --- Position ----------

type TransformComponent struct {
    Position *math.Vector2
    Rotation *math.Vector2
    Scale    *math.Vector2
}

var TransformComponentId = GetWorld().GetNewId()
func (TransformComponent) ComponentId() ComponentId {
    return TransformComponentId
}

// --- Player -----------

type PlayerComponent struct {}

var PlayerComponentId = GetWorld().GetNewId()
func (PlayerComponent) ComponentId() ComponentId {
    return PlayerComponentId
}
