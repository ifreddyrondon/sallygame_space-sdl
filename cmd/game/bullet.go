package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/collider"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/updater"
)

const (
	bulletSpeed           = 15
	bulletCollisionRadius = 8
)

func newBullet(renderer *sdl.Renderer) (*model.Element, error) {
	return model.NewElement(
		model.WithElemCircleCollision(bulletCollisionRadius),
		drawer.WithSpriteRenderer(renderer, fmt.Sprintf("%s/player_bullet.bmp", assetsDir)),
		updater.WithBulletMover(bulletSpeed),
		collider.WithDeactivateCollider(),
	)
}
