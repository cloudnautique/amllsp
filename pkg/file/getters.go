package file

import (
	"os"
	"path/filepath"

	"cuelang.org/go/cue/ast"
)

func (a *Acornfile) Path() string {
	return a.path
}

func (a *Acornfile) Content() *ast.File {
	return a.astContent
}

func (a *Acornfile) Filename() string {
	return filepath.Base(a.path)
}

func (a *Acornfile) StringContent() (string, error) {
	bytes, err := os.ReadFile(a.Path())
	return string(bytes), err
}
