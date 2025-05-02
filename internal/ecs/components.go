package ecs

import "github.com/cxncxl/gogame/internal/math"

// --- Shared ------------

var emptyComponents = map[ComponentId]Component {
    RenderComponentId: RenderComponent{},
    PositionComponentId: PositionComponent{},
    PlayerComponentId: PlayerComponent{},
}

// --- Render ------------

type RenderComponent struct {
    Color [3]uint8
}

var RenderComponentId = GetWorld().GetNewId()
func (RenderComponent) ComponentId() ComponentId {
    return RenderComponentId
}

// --- Position ----------

type PositionComponent struct {
    Position *math.Vector2
}

var PositionComponentId = GetWorld().GetNewId()
func (PositionComponent) ComponentId() ComponentId {
    return PositionComponentId
}

// --- Player -----------

type PlayerComponent struct {}

var PlayerComponentId = GetWorld().GetNewId()
func (PlayerComponent) ComponentId() ComponentId {
    return PlayerComponentId
}
