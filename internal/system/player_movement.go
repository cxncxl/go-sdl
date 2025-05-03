package system

import (
	"log"
	"time"

	"github.com/cxncxl/gogame/internal/ecs"
	"github.com/cxncxl/gogame/internal/math"
	"github.com/cxncxl/gogame/internal/service"
	"github.com/veandco/go-sdl2/sdl"
)

func PlayerMovementSystem(w *ecs.World, dt time.Duration) {
    players := w.Query(ecs.PlayerComponentId)
    if players == nil || len(players) > 1 {
        log.Panic("No player or more than one player")
    }

    player := players[0]

    keysPressed := service.Input().KeyPressed()
    if len(keysPressed) == 0 {
        return
    }

    playerComponents, err := w.GetEntityComponents(player)
    if err != nil {
        log.Println(err)
        return
    }

    playerPosition := playerComponents[ecs.TransformComponentId].(ecs.TransformComponent)

    movement := math.Vector2{}

    for _, keyPressed := range keysPressed {
        switch int(keyPressed) {
        case sdl.K_w:
            movement.Add(north)
        case sdl.K_s:
            movement.Add(south)
        case sdl.K_a:
            movement.Add(west)
        case sdl.K_d:
            movement.Add(east)
        }
    }

    // TODO: pointer chaos
    movementNormalized := movement.Normalized()
    m := &movementNormalized
    m.MulScalar(PlayerSpeed)
    m.MulScalar(float64(dt.Milliseconds()))
    playerPosition.Position.Add(*m);
}

var north = math.Vector2{ X: 0, Y: -1 }
var south = math.Vector2{ X: 0, Y: 1 }
var east = math.Vector2{ X: 1, Y: 0 }
var west = math.Vector2{ X: -1, Y: 0 }

// TODO: this can be variable and should be calculated based on player's
// stats, items, whatever
const PlayerSpeed = 0.3

