package drawer

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
)

type SpriteRenderer struct {
	container   *model.Element
	tex         *sdl.Texture
	srcRect     *sdl.Rect
	centerPoint *sdl.Point
}

func WithSpriteRenderer(renderer *sdl.Renderer, filename string) model.ElemOptFunc {
	builderFn := func(elem *model.Element) (model.Drawer, error) {
		comp, err := NewSpriteRenderer(elem, renderer, filename)
		if err != nil {
			return nil, err
		}
		return comp, err
	}
	return model.WithElemDrawerFn(builderFn)
}

func NewSpriteRenderer(container *model.Element, renderer *sdl.Renderer, filename string) (*SpriteRenderer, error) {
	tex, err := textureFromBMP(renderer, filename)
	if err != nil {
		return nil, err
	}
	_, _, w, h, err := tex.Query()
	if err != nil {
		return nil, err
	}
	return &SpriteRenderer{
		container:   container,
		tex:         tex,
		srcRect:     &sdl.Rect{X: 0, Y: 0, W: w, H: h},
		centerPoint: &sdl.Point{X: w / 2, Y: h / 2},
	}, nil
}

func (sr *SpriteRenderer) OnDraw(renderer *sdl.Renderer) error {
	return drawTexture(sr.tex, sr.container.Position, sr.container.Rotation, renderer)
}

func (sr *SpriteRenderer) SrcRect() *sdl.Rect {
	return sr.srcRect
}
