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
	builderFn := func(elem *model.Element) (model.Updater, error) {
		return &BulletMover{
			container: elem,
			speed:     speed,
		}, nil
	}
	return model.WithElemUpdaterFn(builderFn)
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
