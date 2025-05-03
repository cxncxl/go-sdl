package system

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/cxncxl/gogame/internal/ecs"
	"github.com/cxncxl/gogame/internal/service"
	"github.com/veandco/go-sdl2/sdl"
)

func RenderSystem(w *ecs.World, dt time.Duration) {
    ren := service.Renderer()
    if ren == nil {
        return // todo: log error
    }
    
    ren.SetDrawColor(0x00, 0x00, 0x00, 0xff)
    ren.Clear()

    BoxRenderSystem(w, dt)
    SpriteRenderSystem(w, dt)

    ren.Present()
}

func BoxRenderSystem(w *ecs.World, dt time.Duration) {
    ren := service.Renderer()
    if ren == nil {
        return // todo: log error
    }

    renderables := w.Query(ecs.BoxRendererComponentId, ecs.TransformComponentId)

    ren.SetDrawColor(0x00, 0x00, 0x00, 0xff)

    for _, r := range renderables {
        // todo: culling
        comps, _ := w.GetEntityComponents(r)
        pos, exists := comps[ecs.TransformComponentId]
        if !exists {
            // todo: handle
            fmt.Println("No position comp on entity", r.String())
            continue
        }

        renc, exists := comps[ecs.BoxRendererComponentId]
        if !exists {
            fmt.Println("No render comp on entity", r.String())
            continue
        }

        cpos := pos.(ecs.TransformComponent)
        cren := renc.(ecs.BoxRendererComponent)

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
}

func SpriteRenderSystem(w *ecs.World, dt time.Duration) {
    ren := service.Renderer()
    if ren == nil {
        return
    }

    renderables := w.Query(
        ecs.SpriteRendererComponentId, ecs.TransformComponentId,
    )

    for _, r := range renderables {
        // todo: culling
        comps, _ := w.GetEntityComponents(r)
        pos, exists := comps[ecs.TransformComponentId]
        if !exists {
            // todo: handle
            fmt.Println("No position comp on entity", r.String())
            continue
        }

        renc, exists := comps[ecs.SpriteRendererComponentId]
        if !exists {
            fmt.Println("No render comp on entity", r.String())
            continue
        }

        cpos := pos.(ecs.TransformComponent)
        cren := renc.(ecs.SpriteRendererComponent)

        sprite, err := service.AssetsManager().GetSprite(cren.SpritePath)
        if err != nil {
            if errors.Is(err, service.AssetDoesNotExistError) {
                slog.Error("Sprite", cren.SpritePath, "not loaded")
            } else {
                slog.Error("Unknown error getting sprite" + cren.SpritePath)
                slog.Error(err.Error())
            }

            continue
        }

        box := sdl.Rect{
            X: int32(cpos.Position.X),
            Y: int32(cpos.Position.Y),
            W: 32,
            H: 32,
        }

        // note for AnimationRenderer: second argument crops input texture
        err = ren.Copy(sprite, nil, &box)
    }
}
