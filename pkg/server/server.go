package server

import (
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"
)

type LSP struct {
	handler *Handler

	server *server.Server

	log logging.Logger

	debug bool
}

const (
	lsName = "amllsp"
)

var version string = "0.0.0"

func New() (*LSP, error) {
	baseLog := logging.GetLogger(lsName)

	handler, err := NewHandler(baseLog)
	if err != nil {
		return nil, err
	}

	server := server.NewServer(handler.Handler(), lsName, false)

	return &LSP{
		handler: handler,
		log:     baseLog,
		server:  server,
		debug:   false,
	}, nil
}

func (s *LSP) Run() error {
	s.log.Info("Run server stdio")

	return s.server.RunStdio()
}
