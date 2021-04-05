package updater

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/screen"
)

type KeyboardMover struct {
	container *model.Element
	speed     float64

	sr *drawer.SpriteRenderer
}

func WithKeyboardMover(speed float64) model.ElemOptFunc {
	builderFn := func(elem *model.Element) (model.Updater, error) {
		sr, err := elem.Drawer(&drawer.SpriteRenderer{})
		if err != nil {
			return nil, err
		}
		comp := &KeyboardMover{
			container: elem,
			speed:     speed,
			sr:        sr.(*drawer.SpriteRenderer),
		}
		return comp, nil
	}
	return model.WithElemUpdaterFn(builderFn)
}

func (km KeyboardMover) OnUpdate(delta float64) error {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		if km.container.Position.X-(float64(km.sr.SrcRect().W)/2.0) > 0 {
			km.container.Position.X -= km.speed * (delta)
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if km.container.Position.X+(float64(km.sr.SrcRect().W)/2.0) < float64(screen.Width()) {
			km.container.Position.X += km.speed * (delta)
		}
	}
	return nil
}
