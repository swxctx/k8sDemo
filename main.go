package main

import (
	"daywork/xl/k8Demo/api"
	td "github.com/swxctx/malatd"
)

func main() {
	// Gen Time: 2024-11-18 21:31:36
	srv := td.NewServer(cfg.SrvConfig)
	api.Route(srv, "/k8_demo")
	srv.Run()
}
