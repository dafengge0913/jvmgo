package classpath

import (
	"strings"
	"errors"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (entry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range entry {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entry CompositeEntry) String() string {
	strs := make([]string, len(entry))
	for i, e := range entry {
		strs[i] = e.String()
	}
	return strings.Join(strs, pathListSeparator)
}
