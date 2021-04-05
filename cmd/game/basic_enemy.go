package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/collider"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
)

const (
	basicEnemySize            = 105
	basicEnemyDrawAngle       = 180.0
	basicEnemyCollisionRadius = 38
)

func newBasicEnemy(renderer *sdl.Renderer, position model.Vector) (*model.Element, error) {
	idleSequence, err := drawer.NewAnimationSequence(
		renderer,
		fmt.Sprintf("%s/basic_enemy/idle", assetsDir),
		5, true)
	if err != nil {
		return nil, fmt.Errorf("error creating idle sequence. %w", err)
	}
	destroySequence, err := drawer.NewAnimationSequence(
		renderer,
		fmt.Sprintf("%s/basic_enemy/destroy", assetsDir),
		15, false)
	if err != nil {
		return nil, fmt.Errorf("error creating destroy sequence. %w", err)
	}
	sequences := map[string]*drawer.AnimationSequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}
	return model.NewElement(
		model.WithElemRotation(basicEnemyDrawAngle),
		model.WithElemPosition(position), // bottom center
		model.WithElemActive(),
		model.WithElemCircleCollision(basicEnemyCollisionRadius),
		drawer.WithAnimatorRenderer(sequences, "idle"),
		model.WithElemColliderFn(func(elem *model.Element) (model.Collider, error) {
			comp, err := newBasicEnemyCollider(elem)
			if err != nil {
				return nil, err
			}
			return comp, nil
		}),
	)
}

type basicEnemyCollider struct {
	parent   *collider.DeactivateCollider
	animator *drawer.AnimatorRenderer
}

func newBasicEnemyCollider(elem *model.Element) (*basicEnemyCollider, error) {
	animator, err := elem.Drawer(&drawer.AnimatorRenderer{})
	if err != nil {
		return nil, fmt.Errorf("error getting the animator drawer from element")
	}
	return &basicEnemyCollider{
		parent:   collider.NewDeactivateCollider(elem),
		animator: animator.(*drawer.AnimatorRenderer),
	}, nil
}

func (b *basicEnemyCollider) OnCollision(with *model.Element) error {
	b.animator.SetCurrent("destroy")
	return nil
}
