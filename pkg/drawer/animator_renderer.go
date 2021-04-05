package drawer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
)

var errMissingDefaultSequence = errors.New("missing default sequence")

func WithAnimatorRenderer(sequences map[string]*AnimationSequence, defaultSequence string) model.ElemOptFunc {
	builderFn := func(elem *model.Element) (model.Drawer, error) {
		comp, err := NewAnimator(elem, sequences, defaultSequence)
		if err != nil {
			return nil, err
		}
		return comp, err
	}
	return model.WithElemDrawerFn(builderFn)
}

type AnimatorRenderer struct {
	container *model.Element
	sequences map[string]*AnimationSequence
	current   string
	lastFrame time.Time
}

func NewAnimator(container *model.Element, sequences map[string]*AnimationSequence, defaultSequence string) (*AnimatorRenderer, error) {
	if _, ok := sequences[defaultSequence]; !ok {
		return nil, errMissingDefaultSequence
	}

	return &AnimatorRenderer{
		container: container,
		sequences: sequences,
		current:   defaultSequence,
		lastFrame: time.Now(),
	}, nil
}

func (a *AnimatorRenderer) OnDraw(renderer *sdl.Renderer) error {
	seq := a.sequences[a.current]
	if time.Since(a.lastFrame) >= time.Duration(float64(time.Second)/seq.sampleRate) {
		seq.nextFrame()
		a.lastFrame = time.Now()
	}

	tex := seq.texture()
	return drawTexture(tex, a.container.Position, a.container.Rotation, renderer)
}

func (a *AnimatorRenderer) SetCurrent(current string) {
	a.current = current
	a.lastFrame = time.Now()
}

type AnimationSequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64 // number of times the frame should be incremented per second
	loop       bool
}

func NewAnimationSequence(renderer *sdl.Renderer, filepath string, sampleRate float64, loop bool) (*AnimationSequence, error) {
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading dir %s. Err %w", filepath, err)
	}
	var textures []*sdl.Texture
	for _, f := range files {
		filename := path.Join(filepath, f.Name())
		tex, err := textureFromBMP(renderer, filename)
		if err != nil {
			return nil, err
		}
		textures = append(textures, tex)
	}
	return &AnimationSequence{
		textures:   textures,
		frame:      0,
		sampleRate: sampleRate,
		loop:       loop,
	}, err
}

func (s *AnimationSequence) texture() *sdl.Texture {
	return s.textures[s.frame]
}

func (s *AnimationSequence) nextFrame() {
	if s.frame == len(s.textures)-1 {
		if s.loop {
			s.frame = 0
		}
	} else {
		s.frame++
	}
}
