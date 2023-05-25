package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	shutdownTimeout = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 5 * time.Second
)

// Server provides a http server implementation.
type Server struct {
	ip       string
	port     string
	listener net.Listener
}

// New return new Server and error
func New(port string) (*Server, error) {
	addr := fmt.Sprintf(":" + port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", addr, err)
	}

	return &Server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}

func (s *Server) serveHTTP(ctx context.Context, srv *http.Server) error {
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()
		shutdownCtx, done := context.WithTimeout(context.Background(), shutdownTimeout)
		defer done()
		errCh <- srv.Shutdown(shutdownCtx)
	}()

	// run the server. This will block until the provided context is closed.
	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	// return any errors that happened during shutdown.
	if err := <-errCh; err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	return nil
}

// ServeHTTPHandler starts the server and blocks until the provided context is closed.
// When the provided context is closed, the server is gracefully stopped with a
// timeout of 5 seconds.
func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.serveHTTP(ctx, &http.Server{
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      handler})
}
