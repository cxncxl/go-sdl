package ecs

import "github.com/cxncxl/gogame/internal/math"

// --- Shared ------------

var emptyComponents = map[ComponentId]Component {
    RenderComponentId: RenderComponent{ Color: [3]uint8{ 0, 0, 0 } },
    PositionComponentId: PositionComponent{ Position: math.Vector2{ X: 0, Y: 0 } },
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
    Position math.Vector2
}

var PositionComponentId = GetWorld().GetNewId()
func (PositionComponent) ComponentId() ComponentId {
    return PositionComponentId
}
