package updater

import (
	"math"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/screen"
)

type BulletMover struct {
	container *model.Element
	speed     float64
	delta     *float64
}

func WithBulletMover(speed float64, delta *float64) model.ElemOptFunc {
	return model.WithElemUpdaterFn(func(elem *model.Element) (model.Updater, error) {
		return NewBulletMover(elem, speed, delta), nil
	})
}

func NewBulletMover(container *model.Element, speed float64, delta *float64) *BulletMover {
	return &BulletMover{
		container: container,
		speed:     speed,
		delta:     delta,
	}
}

func (b BulletMover) OnUpdate() error {
	b.container.Position.X += b.speed * (*b.delta) * math.Cos(b.container.Rotation)
	b.container.Position.Y += b.speed * (*b.delta) * math.Sin(b.container.Rotation)

	if b.container.Position.X > float64(screen.Width()) || b.container.Position.X < 0 ||
		b.container.Position.Y > float64(screen.Height()) || b.container.Position.Y < 0 {

		b.container.Active = false
	}
	return nil
}
