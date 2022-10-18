package server

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) documentHover(_ *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	h.log.Infof("pos: %x, %x", params.Position.Line, params.Position.Character)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}

	af := h.workspace.GetFile(_uri.Filename())

	markdown, err := af.GetDocDefinitionMarkDown(
		uintToInt(params.Position.Line),
		uintToInt(params.Position.Character))
	if err != nil {
		return nil, err
	}

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: markdown,
		},
	}, nil
}
