package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	baseServer *http.Server
}

func (srv *Server) Run(port string, handler http.Handler) error {
	srv.baseServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 8 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return srv.baseServer.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.baseServer.Shutdown(ctx)
}
