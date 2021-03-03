package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/collider"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/updater"
)

const (
	bulletSpeed           = 15
	bulletCollisionRadius = 8
)

func newBullet(renderer *sdl.Renderer) (*model.Element, error) {
	return model.NewElement(
		model.WithElemCircleCollision(bulletCollisionRadius),
		drawer.WithSpriteRenderer(renderer, fmt.Sprintf("%s/player_bullet.bmp", assetsDir)),
		updater.WithBulletMover(bulletSpeed, &delta),
		collider.WithDeactivateCollider(),
	)
}
