package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absPath}
}

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(entry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, entry, err
}

func (entry *DirEntry) String() string {
	return entry.absDir
}
