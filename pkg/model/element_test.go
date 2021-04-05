package model_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ifreddyrondon/sallygames_space-sdl/pkg/model"
)

type mockUpdater struct{}

func (e mockUpdater) OnUpdate() error { return nil }

type printUpdater struct{}

func (e printUpdater) OnUpdate() error {
	fmt.Print("update")
	return nil
}

func TestShouldErrWhenCreateElemWithTwoComponentsWithTheSameType(t *testing.T) {
	e, err := model.NewElement(model.WithElemUpdater(mockUpdater{}), model.WithElemUpdater(mockUpdater{}))
	require.Nil(t, e)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "unable to use two or more updaters of the same type.")
	require.Contains(t, err.Error(), ".mockUpdater")
}

func TestShouldErrWhenGetComponentWithNoComponentsOfTheGivenType(t *testing.T) {
	e, err := model.NewElement(model.WithElemUpdater(mockUpdater{}))
	require.Nil(t, err)
	require.NotNil(t, e)
	comp, err := e.Updater(&printUpdater{})
	require.Nil(t, comp)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "missing updater with type:")
	require.Contains(t, err.Error(), ".printUpdater")
}

func TestShouldBeOK(t *testing.T) {
	e, err := model.NewElement(
		model.WithElemUpdater(mockUpdater{}),
		model.WithElemPosition(model.Vector{X: 1, Y: 2}),
		model.WithElemRotation(30),
		model.WithElemActive(),
	)
	require.Nil(t, err)
	require.NotNil(t, e)
	require.True(t, e.Active)
	require.Equal(t, 1.0, e.Position.X)
	require.Equal(t, 2.0, e.Position.Y)
	require.Equal(t, 30.0, e.Rotation)
}
