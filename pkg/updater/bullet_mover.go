package updater

import (
	"math"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/screen"
)

type BulletMover struct {
	container *model.Element
	speed     float64
}

func WithBulletMover(speed float64) model.ElemOptFunc {
	return model.WithElemUpdaterFn(func(elem *model.Element) (model.Updater, error) {
		return NewBulletMover(elem, speed), nil
	})
}

func NewBulletMover(container *model.Element, speed float64) *BulletMover {
	return &BulletMover{
		container: container,
		speed:     speed,
	}
}

func (b BulletMover) OnUpdate(delta float64) error {
	b.container.Position.X += b.speed * (delta) * math.Cos(b.container.Rotation)
	b.container.Position.Y += b.speed * (delta) * math.Sin(b.container.Rotation)

	if b.container.Position.X > float64(screen.Width()) || b.container.Position.X < 0 ||
		b.container.Position.Y > float64(screen.Height()) || b.container.Position.Y < 0 {

		b.container.Active = false
	}
	return nil
}
