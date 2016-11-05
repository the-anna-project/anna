package spec

import (
	"os"
)

// FileSystem provides certain file system implementations for abstraction and
// testing reasons.
type FileSystem interface {
	// TODO this should may be Metadata ???
	//Object

	// ReadFile is aligned to ioutil.ReadFile.
	ReadFile(filename string) ([]byte, error)

	// WriteFile is aligned to ioutil.WriteFile.
	WriteFile(filename string, bytes []byte, perm os.FileMode) error
}
