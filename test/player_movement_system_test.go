package test

import (
	"testing"
	"time"

	"github.com/cxncxl/gogame/internal/ecs"
	"github.com/cxncxl/gogame/internal/math"
	"github.com/cxncxl/gogame/internal/service"
	"github.com/cxncxl/gogame/internal/system"
	"github.com/veandco/go-sdl2/sdl"
)

func setupTest() (*ecs.World, *math.Vector2) {
    w := ecs.GetWorld()
    playerRec := w.NewEntity(
        ecs.NewType(
            ecs.PlayerComponentId,
            ecs.TransformComponentId,
        ),
    )

    playerPos := math.Vector2{ X: 0, Y: 0 }

    w.SetEntityComponent(playerRec.Entity, ecs.TransformComponent{
        Position: &playerPos,
    })

    service.InitInput()

    return w, &playerPos
}

func TestPlayerMovement(t *testing.T) {
    w, position := setupTest()

    event := sdl.KeyboardEvent{
        Type: sdl.KEYDOWN,
        Keysym: sdl.Keysym{
            Sym: sdl.K_w,
        },
    }

    service.Input().HandleEvent(&event)

    system.PlayerMovementSystem(w, time.Microsecond)

    if position.Y != -1 * system.PlayerSpeed {
        t.Errorf(
            "Movement not handling W button. Expected position: %v, actual: %v\n",
            math.Vector2{ Y: -1 * system.PlayerSpeed },
            position,
        )
    }

    event = sdl.KeyboardEvent{
        Type: sdl.KEYUP,
        Keysym: sdl.Keysym{
            Sym: sdl.K_w,
        },
    }

    service.Input().HandleEvent(&event)

    system.PlayerMovementSystem(w, time.Microsecond)
    
    if position.Y != -1 * system.PlayerSpeed {
        t.Errorf(
            "Movement not handling button release. Expected position: %v, actual: %v\n",
            math.Vector2{ Y: -1 * system.PlayerSpeed },
            position,
        )
    }
}
