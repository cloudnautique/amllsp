package main

import (
	"log"

	"github.com/acorn-io/amllsp/pkg/server"

	"github.com/tliron/kutil/logging"
	_ "github.com/tliron/kutil/logging/simple"
)

func main() {
	logging.Configure(1, nil)
	srv, err := server.New()
	if err != nil {
		log.Fatalf("could not init AML Language server, %s", err)
	}
	srv.Run()
}
