package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/cxncxl/gogame/internal/ecs"
	"github.com/cxncxl/gogame/internal/math"
	"github.com/cxncxl/gogame/internal/service"
)

func main() {
    sdl.Init(sdl.INIT_EVERYTHING)
    defer sdl.Quit()

    service.InitRenderer()
    defer service.DeinitRenderer()

    world := ecs.GetWorld()

    for i := range 255 {
        rec := world.NewEntity(ecs.NewType(ecs.PositionComponentId, ecs.RenderComponentId))
        world.SetEntityComponent(rec.Entity, ecs.RenderComponent{
            Color: [3]uint8{ 
                0xff - uint8(i),
                0x00 + (uint8(i) * 2),
                0xff,
            },
        })
        world.SetEntityComponent(rec.Entity, ecs.PositionComponent{
            Position: math.Vector2{
                X: float64(50 + (uint(i) * 20)),
                Y: float64(50 + (uint(i) * 10)),
            },
        })
    }

    running := true
    prevTime := time.Now()
    dt := time.Millisecond 
    for running == true {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                running = false
                break
            }
        }

        for _, s := range ecs.Systems {
            s(world, dt)
        }

        t := time.Now()
        dt = t.Sub(prevTime)

        fmt.Println("Peformace ===========================")
        fmt.Println("\tprevTime:", prevTime.Format(time.RFC3339Nano))
        fmt.Println("\tcurrent time:", t.Format(time.RFC3339Nano))
        fmt.Println("\tdt:", dt)
        fmt.Println("\tfps:", 1_000_000_000 / dt.Nanoseconds())
        fmt.Println("=====================================")

        prevTime = t
    }
}
