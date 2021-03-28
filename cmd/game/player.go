package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/drawer"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/updater"
)

const (
	playerSpeed        = 5
	playerSize         = 105
	playerShotCoolDown = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer, bullets []*model.Element) (*model.Element, error) {
	return model.NewElement(
		model.WithElemPosition(model.Vector{X: ScreenWidth / 2.0, Y: ScreenHeight - playerSize/2.0}), // bottom center
		model.WithElemActive(),
		drawer.WithSpriteRenderer(renderer, fmt.Sprintf("%s/player.bmp", assetsDir)),
		updater.WithKeyboardMover(playerSpeed),
		updater.WithKeyboardShooter(model.NewPool(bullets), playerShotCoolDown),
	)
}
