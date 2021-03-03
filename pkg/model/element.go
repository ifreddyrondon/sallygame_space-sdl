package model

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	errEmptyUpdater  = fmt.Errorf("a non nil updater is required")
	errEmptyDrawer   = fmt.Errorf("a non nil drawer is required")
	errEmptyCollider = fmt.Errorf("a non nil collider is required")
)

func WithElemPosition(position Vector) ElemOptFunc {
	return func(elem *Element) error {
		elem.Position = position
		return nil
	}
}

func WithElemRotation(rotation float64) ElemOptFunc {
	return func(elem *Element) error {
		elem.Rotation = rotation
		return nil
	}
}

func WithElemActive() ElemOptFunc {
	return func(elem *Element) error {
		elem.Active = true
		return nil
	}
}

func WithElemCircleCollision(radius float64) ElemOptFunc {
	return func(elem *Element) error {
		elem.collisions = append(elem.collisions, newCircleCollision(elem, radius))
		return nil
	}
}

func WithElemUpdater(comp Updater) ElemOptFunc {
	return func(elem *Element) error {
		return checkUpdater(elem, comp)
	}
}

func WithElemUpdaterFn(fn func(elem *Element) (Updater, error)) ElemOptFunc {
	return func(elem *Element) error {
		comp, err := fn(elem)
		if err != nil {
			return err
		}
		return checkUpdater(elem, comp)
	}
}

func checkUpdater(elem *Element, comp Updater) error {
	if comp == nil {
		return errEmptyUpdater
	}
	for i := 0; i < len(elem.updaters); i++ {
		if reflect.TypeOf(comp) == reflect.TypeOf(elem.updaters[i]) {
			return fmt.Errorf("unable to use two or more updaters of the same type. Type: %v", reflect.TypeOf(comp))
		}
	}
	elem.updaters = append(elem.updaters, comp)
	return nil
}

func WithElemDrawer(comp Drawer) ElemOptFunc {
	return func(elem *Element) error {
		return checkDrawer(elem, comp)
	}
}

func WithElemDrawerFn(fn func(elem *Element) (Drawer, error)) ElemOptFunc {
	return func(elem *Element) error {
		comp, err := fn(elem)
		if err != nil {
			return err
		}
		return checkDrawer(elem, comp)
	}
}

func checkDrawer(elem *Element, comp Drawer) error {
	if comp == nil {
		return errEmptyDrawer
	}
	for i := 0; i < len(elem.drawers); i++ {
		if reflect.TypeOf(comp) == reflect.TypeOf(elem.drawers[i]) {
			return fmt.Errorf("unable to use two or more drawers of the same type. Type: %v", reflect.TypeOf(comp))
		}
	}
	elem.drawers = append(elem.drawers, comp)
	return nil
}

func WithElemCollider(comp Collider) ElemOptFunc {
	return func(elem *Element) error {
		return checkCollider(elem, comp)
	}
}

func WithElemColliderrFn(fn func(elem *Element) (Collider, error)) ElemOptFunc {
	return func(elem *Element) error {
		comp, err := fn(elem)
		if err != nil {
			return err
		}
		return checkCollider(elem, comp)
	}
}

func checkCollider(elem *Element, comp Collider) error {
	if comp == nil {
		return errEmptyCollider
	}
	for i := 0; i < len(elem.colliders); i++ {
		if reflect.TypeOf(comp) == reflect.TypeOf(elem.colliders[i]) {
			return fmt.Errorf("unable to use two or more colliders of the same type. Type: %v", reflect.TypeOf(comp))
		}
	}
	elem.colliders = append(elem.colliders, comp)
	return nil
}

type ElemOptFunc func(opts *Element) error

type Updater interface {
	OnUpdate() error
}

type Drawer interface {
	OnDraw(renderer *sdl.Renderer) error
}

type Collider interface {
	OnCollision(with *Element) error
}

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	collisions []circleCollision

	updaters  []Updater
	drawers   []Drawer
	colliders []Collider
}

func NewElement(opts ...ElemOptFunc) (*Element, error) {
	var elem Element
	for _, opt := range opts {
		if err := opt(&elem); err != nil {
			return nil, err
		}
	}
	return &elem, nil
}

func (e *Element) Drawer(withType Drawer) (Drawer, error) {
	typ := reflect.TypeOf(withType)
	for _, comp := range e.drawers {
		if typ == reflect.TypeOf(comp) {
			return comp, nil
		}
	}
	return nil, fmt.Errorf("missing drawer with type: %v", typ)
}

func (e *Element) Updater(withType Updater) (Updater, error) {
	typ := reflect.TypeOf(withType)
	for _, comp := range e.updaters {
		if typ == reflect.TypeOf(comp) {
			return comp, nil
		}
	}
	return nil, fmt.Errorf("missing updater with type: %v", typ)
}

func (e *Element) Draw(renderer *sdl.Renderer) error {
	if !e.Active {
		return nil
	}
	for _, comp := range e.drawers {
		if err := comp.OnDraw(renderer); err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) Update() error {
	if !e.Active {
		return nil
	}
	for _, comp := range e.updaters {
		if err := comp.OnUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) Collision(with *Element) error {
	if !e.Active {
		return nil
	}
	for _, comp := range e.colliders {
		if err := comp.OnCollision(with); err != nil {
			return err
		}
	}
	return nil
}
