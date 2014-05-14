package watcher

import (
	"code.google.com/p/go.exp/fsnotify"
	"os"
	"path/filepath"
)

type Watcher struct {
	*fsnotify.Watcher
}

func NewWatcher(path string, recursive bool) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	rw := &Watcher{Watcher: watcher}

	if recursive {
		if err := rw.addSubFolders(path); err != nil {
			return nil, err
		}
	} else {
		rw.Watch(path)
	}

	return rw, nil
}

func (rw *Watcher) addSubFolders(path string) error {
	return filepath.Walk(path, func(subdir string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			name := info.Name()
			hidden := filepath.HasPrefix(name, ".") && name != "." && name != ".."
			if hidden {
				return filepath.SkipDir
			}

			rw.Watcher.Watch(subdir)
		}

		return nil
	})
}
