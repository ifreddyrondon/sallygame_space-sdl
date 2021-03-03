package model

import (
	"math"
)

type circleCollision struct {
	container *Element
	radius    float64
}

func newCircleCollision(container *Element, radius float64) circleCollision {
	return circleCollision{
		container: container,
		radius:    radius,
	}
}

func collides(c1, c2 circleCollision) bool {
	dist := math.Sqrt(
		math.Pow(c2.container.Position.X-c1.container.Position.X, 2) +
			math.Pow(c2.container.Position.Y-c1.container.Position.Y, 2),
	)
	return dist <= c1.radius+c2.radius
}

func CheckCollisions(elements []*Element) error {
	for i := 0; i < len(elements); i++ {
		e1 := elements[i]
		if !e1.Active {
			continue
		}
		for j := i + 1; j < len(elements); j++ {
			e2 := elements[j]
			if !e2.Active {
				continue
			}
			for _, c1 := range e1.collisions {
				for _, c2 := range e2.collisions {
					if !collides(c1, c2) {
						continue
					}
					if err := e1.Collision(e2); err != nil {
						return err
					}
					if err := e2.Collision(e1); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
