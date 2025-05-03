package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/cxncxl/gogame/internal/ecs"
	"github.com/cxncxl/gogame/internal/math"
	"github.com/cxncxl/gogame/internal/service"
	"github.com/cxncxl/gogame/internal/system"
)

func main() {
    sdl.Init(sdl.INIT_EVERYTHING)
    defer sdl.Quit()

    service.InitRenderer()
    defer service.DeinitRenderer()

    service.InitInput()
    defer service.DeinitInput()

    service.InitAssetsManager()
    defer service.DeinitAssetsManager()

    world := ecs.GetWorld()

    player := world.NewEntity(
        ecs.NewType(
            ecs.PlayerComponentId,
            ecs.TransformComponentId,
            ecs.SpriteRendererComponentId,
        ),
    )
    world.SetEntityComponent(player.Entity, ecs.PlayerComponent{})
    world.SetEntityComponent(player.Entity, ecs.TransformComponent{
        Position: &math.Vector2{ X: 0, Y: 0 },
    })
    world.SetEntityComponent(player.Entity, ecs.SpriteRendererComponent{
        SpritePath: "image.png",
    })
    running := true
    dt := time.Millisecond 
    t := time.Now()
    for running == true {
        clearFrame()

        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                running = false
                break
            case *sdl.KeyboardEvent:
                e := event.(*sdl.KeyboardEvent)
                service.Input().HandleEvent(e)
            }
        }

        for _, s := range system.Systems {
            s(world, dt)
        }

        dt, t = updateTime(t)
    }
}

func clearFrame() {}

func updateTime(prevTime time.Time) (time.Duration, time.Time) {
    t := time.Now()
    dt := t.Sub(prevTime)

    fmt.Println("Peformace ===========================")
    fmt.Println("\tprevTime:", prevTime.Format(time.RFC3339Nano))
    fmt.Println("\tcurrent time:", t.Format(time.RFC3339Nano))
    fmt.Println("\tdt:", dt)
    fmt.Println("\tfps:", 1_000_000_000 / dt.Nanoseconds())
    fmt.Println("=====================================")

    return dt, t
}
