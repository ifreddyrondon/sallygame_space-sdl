package main

import (
	"context"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_clue-all/pkg/model"
)

func loop(ctx context.Context, renderer *sdl.Renderer, elems []*model.Element) <-chan error {
	errCh := make(chan error)
	go func() {
		defer close(errCh)
		for {
			select {
			case <-ctx.Done():
				errCh <- nil
				return
			default:
				fmeStartTime := time.Now()
				_ = renderer.Clear()
				for _, elem := range elems {
					if err := elem.Update(delta); err != nil {
						errCh <- fmt.Errorf("error updating element. %w", err)
						return
					}
					if err := elem.Draw(renderer); err != nil {
						errCh <- fmt.Errorf("error drawing element. %w", err)
						return
					}
				}

				if err := model.CheckCollisions(elems); err != nil {
					errCh <- fmt.Errorf("error checking collisions. %w", err)
					return
				}
				renderer.Present()
				delta = time.Since(fmeStartTime).Seconds() * fps
			}
		}
	}()
	return errCh
}
