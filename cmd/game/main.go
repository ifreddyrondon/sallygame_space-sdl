package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
	"github.com/ifreddyrondon/sallygames_clue-all/pkg/scaffolding"
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
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "startup error init sdl. %s\n", err)
		os.Exit(1)
	}

	win, err := sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "startup error init window. %s\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = win.Destroy()
	}()

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "startup error init renderer. %s\n", err)
		os.Exit(1)
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
			_, _ = fmt.Fprintf(os.Stderr, "startup error creating the bullet for palyer. %s\n", err)
			os.Exit(1)
		}
		playerBullets = append(playerBullets, b)
		elements = append(elements, b)
	}

	plr, err := newPlayer(renderer, playerBullets)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "startup error creating the player. %s\n", err)
		os.Exit(1)
	}
	elements = append(elements, plr)

	_ = renderer.SetDrawColor(255, 255, 255, 255)
	for {
		fmeStartTime := time.Now()
		for evt := sdl.PollEvent(); evt != nil; evt = sdl.PollEvent() {
			switch evt.GetType() {
			case sdl.QUIT:
				// TODO: destroy elements
				return
			}
		}

		_ = renderer.Clear()
		for _, elem := range elements {
			if err := elem.Update(delta); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error updating element. Err: %v\n", err)
				return
			}
			if err := elem.Draw(renderer); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error drawing element. Err: %v\n", err)
				return
			}
		}

		if err := model.CheckCollisions(elements); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error checking collisions. Err: %v\n", err)
			return
		}
		renderer.Present()
		delta = time.Since(fmeStartTime).Seconds() * fps
	}
}
