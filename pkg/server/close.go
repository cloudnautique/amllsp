package server

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) didClose(_ *glsp.Context, _ *protocol.DidCloseTextDocumentParams) error {
	return nil
}
