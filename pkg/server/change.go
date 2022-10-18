package server

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) didChange(_ *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	return nil
}
