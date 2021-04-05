package main

import (
	"context"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/scaffolding"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 800
	winTitle     = "Clue"

	fps = 60 // target ticks per seconds
)

var (
	assetsDir = func() string {
		appDir, err := scaffolding.AppDir()
		if err != nil {
			panic(fmt.Errorf("startup error searching appDir. %w", err))
		}
		return fmt.Sprintf("%s/assets", appDir)
	}()

	delta float64
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("startup error init sdl. %w", err)
	}
	defer sdl.Quit()

	win, err := sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("startup error init window. %w", err)
	}
	defer func() {
		_ = win.Destroy()
	}()

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("startup error init renderer. %w", err)
	}
	defer func() {
		_ = renderer.Destroy()
	}()

	var elements []*model.Element
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*ScreenWidth + (basicEnemySize / 2)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2)

			enemy, err := newBasicEnemy(renderer, model.Vector{X: x, Y: y})
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "startup error creating the basic enemy. %s\n", err)
				os.Exit(1)
			}
			elements = append(elements, enemy)
		}
	}
	var playerBullets []*model.Element
	for i := 0; i < 30; i++ {
		b, err := newBullet(renderer)
		if err != nil {
			return fmt.Errorf("startup error creating the bullet for player. %w", err)
		}
		playerBullets = append(playerBullets, b)
		elements = append(elements, b)
	}

	plr, err := newPlayer(renderer, playerBullets)
	if err != nil {
		return fmt.Errorf("startup error creating the player. %w", err)
	}
	elements = append(elements, plr)

	_ = renderer.SetDrawColor(255, 255, 255, 255)

	ctx, cancel := context.WithCancel(context.Background())
	errCh := loop(ctx, renderer, elements)

	for {
		select {
		case err := <-errCh:
			return err
		default:
			evt := sdl.WaitEvent()
			switch evt.GetType() {
			case sdl.QUIT:
				cancel()
			}
		}
	}
}
