package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

func main() {
    sdl.Init(sdl.INIT_EVERYTHING)
    defer sdl.Quit()

    win, err := sdl.CreateWindow(
        "Go Game",
        0, 0,
        800, 600,
        sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE,
    )
    if err != nil {
        panic(err)
    }
    defer win.Destroy()

    running := true
    for running == true {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                running = false
                break
            }
        }
    }
}
