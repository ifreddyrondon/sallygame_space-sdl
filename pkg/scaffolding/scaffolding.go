package scaffolding

import (
	"errors"
	"os"
	"path/filepath"
)

type fsGetter interface {
	Getwd() (dir string, err error)
}

type fsGetterFunc func() (dir string, err error)

func (f fsGetterFunc) Getwd() (dir string, err error) {
	return f()
}

var (
	osFSGetter           = fsGetterFunc(os.Getwd)
	crrFSGetter fsGetter = osFSGetter

	errMissingMod = errors.New("unable to find go.mod, missing app dir")
)

// AppDir returns application path where the app is running.
// it traverse folders upwards from crr until he finds a go.mod
func AppDir() (string, error) {
	for crrPath, _ := crrFSGetter.Getwd(); crrPath != string(filepath.Separator); {
		if _, err := os.Stat(filepath.Join(crrPath, "go.mod")); err == nil {
			return crrPath, nil
		}
		crrPath = filepath.Dir(crrPath)
	}
	return "", errMissingMod
}
