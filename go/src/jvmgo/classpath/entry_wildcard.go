package classpath

import (
	"path/filepath"
	"os"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if isZipEntry(path) {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
