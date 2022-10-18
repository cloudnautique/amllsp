package server

import (
	"fmt"

	"github.com/acorn-io/amllsp/pkg/workspace"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
	"go.lsp.dev/uri"
)

type Handler struct {
	handler   protocol.Handler
	log       logging.Logger
	workspace *workspace.Workspace
}

func NewHandler(logger logging.Logger) (*Handler, error) {
	h := &Handler{
		log: logging.NewScopeLogger(logger, "handler"),
	}

	h.handler = protocol.Handler{
		Initialize:  h.initialize,
		Initialized: h.initialized,
		Shutdown:    h.shutdown,
		SetTrace:    h.setTrace,
		//TextDocumentCompletion: h.completion,
		TextDocumentDidSave:   h.didSave,
		TextDocumentDidOpen:   h.didOpen,
		TextDocumentHover:     h.documentHover,
		TextDocumentDidChange: h.didChange,
		TextDocumentDidClose:  h.didClose,
	}

	return h, nil
}

func (h *Handler) Handler() *protocol.Handler {
	return &h.handler
}
func (h *Handler) initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := h.handler.CreateServerCapabilities()

	//capabilities.CompletionProvider.TriggerCharacters = []string{"."}

	if err := h.initWorkspace(params.WorkspaceFolders, params.RootURI, params.RootPath); err != nil {
		return nil, err
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func (h *Handler) initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func (h *Handler) shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func (h *Handler) setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func (h *Handler) initWorkspace(folders []protocol.WorkspaceFolder, rootURI, rootPath *string) error {
	switch len(folders) {
	case 0:
		var path string
		switch {
		case rootURI != nil:
			path = *rootURI
		case rootPath != nil:
			path = *rootPath
		default:
			return fmt.Errorf("no workspace found")
		}

		_uri, err := uri.Parse(path)
		if err != nil {
			return err
		}

		h.workspace = workspace.New(_uri.Filename(), h.log)

		return nil
	case 1:
		_uri, err := uri.Parse(folders[0].URI)
		if err != nil {
			return err
		}

		h.workspace = workspace.New(_uri.Filename(), h.log)
		return nil
	default:
		return fmt.Errorf("unsupported multiworkspace")
	}
}
