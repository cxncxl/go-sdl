package ecs

// --- Render ------------

type RenderComponent struct {
    Color [3]uint8
}

var RenderComponentId = ComponentId{
    raw: 1, // ???
}

func (RenderComponent) ComponentId() ComponentId {
    return RenderComponentId
}

// --- Position ------------

type PositionComponent struct {
    X uint
    Y uint
}

var PositionComponentId = ComponentId{
    raw: 2,
}

func (PositionComponent) ComponentId() ComponentId {
    return PositionComponentId
}
