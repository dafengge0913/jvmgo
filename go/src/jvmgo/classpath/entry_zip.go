package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(entry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			data, err := ioutil.ReadAll(rc)
			rc.Close()
			if err != nil {
				return nil, nil, err
			}
			return data, entry, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}
