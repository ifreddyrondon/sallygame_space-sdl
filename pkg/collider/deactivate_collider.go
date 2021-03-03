package collider

import (
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
)

type DeactivateCollider struct {
	container *model.Element
}

func WithDeactivateCollider() model.ElemOptFunc {
	return model.WithElemColliderrFn(func(elem *model.Element) (model.Collider, error) {
		return NewDeactivateCollider(elem), nil
	})
}

func NewDeactivateCollider(container *model.Element) *DeactivateCollider {
	return &DeactivateCollider{container: container}
}

func (b *DeactivateCollider) OnCollision(with *model.Element) error {
	b.container.Active = false
	return nil
}
