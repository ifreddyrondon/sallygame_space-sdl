package updater

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
)

type KeyboardShooter struct {
	container *model.Element
	bullets   model.Pool
	coolDown  time.Duration
	lastShot  time.Time
	sr        *drawer.SpriteRenderer
}

func WithKeyboardShooter(bullets model.Pool, coolDown time.Duration) model.ElemOptFunc {
	return model.WithElemUpdaterFn(func(elem *model.Element) (model.Updater, error) {
		return NewKeyboardShooter(elem, bullets, coolDown), nil
	})
}

func NewKeyboardShooter(container *model.Element, bullets model.Pool, coolDown time.Duration) *KeyboardShooter {
	return &KeyboardShooter{
		container: container,
		bullets:   bullets,
		coolDown:  coolDown,
	}
}

func (ks *KeyboardShooter) OnUpdate(delta float64) error {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(ks.lastShot) >= ks.coolDown {
			ks.shot(ks.container.Position.X+25, ks.container.Position.Y-20)
			ks.shot(ks.container.Position.X-25, ks.container.Position.Y-20)
			ks.lastShot = time.Now()
		}
	}
	return nil
}

func (ks *KeyboardShooter) shot(x, y float64) {
	if bul, ok := ks.bullets.Get(); ok {
		bul.Active = true
		bul.Position.X = x
		bul.Position.Y = y
		bul.Rotation = 270 * (math.Pi / 180) // angle to radian
	}
}
