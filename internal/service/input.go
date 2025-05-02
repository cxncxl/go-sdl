package service

import (
	"slices"

	"github.com/veandco/go-sdl2/sdl"
)

var service InputService

func InitInput() {
    service = InputService{
        keysPressed: make([]sdl.Keycode, 0, 100),
    }
}

func DeinitInput() {
}

func Input() *InputService {
    return &service
}

type InputService struct {
    keysPressed []sdl.Keycode
}

func (self *InputService) HandleEvent(event *sdl.KeyboardEvent) {
    idx := slices.Index(self.keysPressed, event.Keysym.Sym)

    if event.GetType() == sdl.KEYUP {
        if idx < 0 {
            return
        }

        self.keysPressed = slices.Delete(self.keysPressed, idx, idx + 1)

        return
    }

    if idx < 0 {
        self.keysPressed = append(self.keysPressed, event.Keysym.Sym)
    }
}

func (self *InputService) Clear() {
    self.keysPressed = []sdl.Keycode{}
}

func (self InputService) KeyPressed() []sdl.Keycode {
    return self.keysPressed
}
