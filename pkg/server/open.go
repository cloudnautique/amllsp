package server

import (
	"github.com/acorn-io/amllsp/pkg/file"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) didOpen(_ *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return err
	}

	a, err := file.New(_uri.Filename(), h.log)
	if err != nil {
		return err
	}

	return h.workspace.AddFile(a)
}
