package drawer

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
)

func textureFromBMP(renderer *sdl.Renderer, filename string) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, fmt.Errorf("startup error loading basic enemy bmp. %w", err)
	}
	defer img.Free()
	return renderer.CreateTextureFromSurface(img)
}

func drawTexture(tex *sdl.Texture, position model.Vector, rotation float64, renderer *sdl.Renderer) error {
	_, _, w, h, err := tex.Query()
	if err != nil {
		return fmt.Errorf("error querying the texture. %w", err)
	}
	center := sdl.Point{X: w / 2, Y: h / 2}
	// converting sprite coordinates to top left of sprite
	x := int32(position.X) - center.X
	y := int32(position.Y) - center.Y

	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: w, H: h},
		&sdl.Rect{X: x, Y: y, W: w, H: h},
		rotation,
		&center,
		sdl.FLIP_NONE,
	)
}
