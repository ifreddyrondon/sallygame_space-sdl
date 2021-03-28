package updater

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/screen"
)

type KeyboardMover struct {
	container *model.Element
	speed     float64

	sr *drawer.SpriteRenderer
}

func WithKeyboardMover(speed float64) model.ElemOptFunc {
	return model.WithElemUpdaterFn(func(elem *model.Element) (model.Updater, error) {
		comp, err := NewKeyboardMover(elem, speed)
		if err != nil {
			return nil, err
		}
		return comp, nil
	})

}

func NewKeyboardMover(container *model.Element, speed float64) (*KeyboardMover, error) {
	sr, err := container.Drawer(&drawer.SpriteRenderer{})
	if err != nil {
		return nil, err
	}
	return &KeyboardMover{
		container: container,
		speed:     speed,
		sr:        sr.(*drawer.SpriteRenderer),
	}, nil
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
