package server

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) didSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	return nil
}
