package ecs

import (
	"fmt"
	"log"
	"time"

	"github.com/cxncxl/gogame/internal/math"
	"github.com/cxncxl/gogame/internal/service"
	"github.com/veandco/go-sdl2/sdl"
)

var Systems = []func(world *World, dt time.Duration){
    RenderSystem,
    PlayerMovementSystem,
};

func RenderSystem(w *World, dt time.Duration) {
    ren := service.Renderer()
    if ren == nil {
        return // todo: log error
    }

    renderables := w.Query(RenderComponentId, PositionComponentId)

    ren.SetDrawColor(0x00, 0x00, 0x00, 0xff)
    ren.Clear()

    for _, r := range renderables {
        // todo: culling
        comps, _ := w.GetEntityComponents(r)
        pos, exists := comps[PositionComponentId]
        if !exists {
            // todo: handle
            fmt.Println("No position comp on entity", r.String())
            continue
        }

        renc, exists := comps[RenderComponentId]
        if !exists {
            fmt.Println("No render comp on entity", r.String())
            continue
        }

        cpos := pos.(PositionComponent)
        cren := renc.(RenderComponent)

        ren.SetDrawColor(cren.Color[0], cren.Color[1], cren.Color[2], 0xff)
        err := ren.FillRect(
            &sdl.Rect{
                // todo: map world coords to screen coords
                X: int32(cpos.Position.X),
                Y: int32(cpos.Position.Y),
                W: 20,
                H: 20,
            },
        )
        if err != nil {
            fmt.Println("Error drawing rect:", err)
        }
    }

    ren.Present()
}

func PlayerMovementSystem(w *World, dt time.Duration) {
    players := w.Query(PlayerComponentId)
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

    playerPosition := playerComponents[PositionComponentId].(PositionComponent)

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
    playerPosition.Position.Add(*m);
}

var north = math.Vector2{ X: 0, Y: -1 }
var south = math.Vector2{ X: 0, Y: 1 }
var east = math.Vector2{ X: 1, Y: 0 }
var west = math.Vector2{ X: -1, Y: 0 }

// TODO: this can be variable and should be calculated based on player's
// stats, items, whatever
const PlayerSpeed = 3
