package httpserver

import (
	"context"
	"github.com/oleglacto/traning_scheduler/internal/app/training_scheduler/httpserver/routes"
	"net/http"
)

type ServerInterface interface {
	Serve() error
	Shutdown(ctx context.Context) error
	GetContext() context.Context
}

type Server struct {
	server  *http.Server
	context context.Context
	cancel  context.CancelFunc
}

func NewServer(addr string, port string) ServerInterface {
	server := &http.Server{Addr: addr + ":" + port, Handler: routes.InitRouter()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	return Server{
		server:  server,
		context: serverCtx,
		cancel:  serverStopCtx,
	}
}

func (s Server) GetContext() context.Context {
	return s.context
}

func (s Server) Serve() error {
	return s.server.ListenAndServe()
}

func (s Server) Shutdown(ctx context.Context) error {
	err := s.server.Shutdown(ctx)

	if err != nil {
		return err
	}

	s.cancel()
	return nil
}
