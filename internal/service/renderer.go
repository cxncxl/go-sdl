package service

import "github.com/veandco/go-sdl2/sdl"

var window *sdl.Window
var renderer *sdl.Renderer

func InitRenderer() {
    if window == nil {
        win, err := sdl.CreateWindow(
            "Go Game",
            0, 0,
            800, 600,
            sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE,
        )
        if err != nil {
            panic(err)
        }

        window = win
        renderer = nil // we want to use new window i guess
    }

    if renderer == nil {
        ren, err := sdl.CreateRenderer(
            window,
            -1, // -1 to initialize the first one supporting the requested flags
            sdl.RENDERER_PRESENTVSYNC,
        )
        if err != nil {
            panic(err)
        }

        renderer = ren
    }
}

func DeinitRenderer() {
    if renderer != nil {
        renderer.Destroy()
    }

    if window != nil {
        window.Destroy()
    }
}

func Renderer() *sdl.Renderer {
    return renderer
}
