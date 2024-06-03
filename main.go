package main

import (
	"net/http"

	"github.com/artemxgod/study/study-uber-fx/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			server.NewHTTPServer,
			server.NewEchoHandler,
			server.NewServeMux,
		),
		fx.Invoke(func(s *http.Server) {}),
	).Run()
}
