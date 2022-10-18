package workspace

import (
	"github.com/acorn-io/amllsp/pkg/file"
	"github.com/tliron/kutil/logging"
)

type Workspace struct {
	path  string
	log   logging.Logger
	files map[string]*file.Acornfile
}

func New(path string, logger logging.Logger) *Workspace {
	return &Workspace{
		path:  path,
		log:   logger,
		files: map[string]*file.Acornfile{},
	}
}

func (wk *Workspace) AddFile(file *file.Acornfile) error {
	wk.files[file.Path()] = file
	return nil
}

func (wk *Workspace) GetFile(path string) *file.Acornfile {
	if val, ok := wk.files[path]; !ok {
		return nil
	} else {
		return val
	}
}
