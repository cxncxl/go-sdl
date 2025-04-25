package ecs

import (
	"fmt"
	"time"

	"github.com/cxncxl/gogame/internal/service"
	"github.com/veandco/go-sdl2/sdl"
)

var Systems = []func(world *World, dt time.Duration){
    RenderSystem,
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
                X: int32(cpos.X),
                Y: int32(cpos.Y),
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
