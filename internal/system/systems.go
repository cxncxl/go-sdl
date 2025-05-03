package system

import (
	"time"

	"github.com/cxncxl/gogame/internal/ecs"
)

type System = func(world *ecs.World, dt time.Duration)

var Systems = []System {
    PlayerMovementSystem,
    RenderSystem,
};
