// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/tylerb/graceful"

	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

// New creates a new server service.
func New() servicespec.Server {
	return &service{}
}

type service struct {
	// Dependencies.

	instrumentation   systemspec.Instrumentation
	serviceCollection servicespec.Collection

	// Settings.

	// httpAddr is the host:port representation based on the golang convention
	// for http.ListenAndServe to serve HTTP traffic.
	httpAddr     string
	bootOnce     sync.Once
	closer       chan struct{}
	httpServer   *graceful.Server
	metadata     map[string]string
	shutdownOnce sync.Once
}

func (s *service) Boot() {
	s.Service().Log().Line("func", "Boot")

	s.bootOnce.Do(func() {
		// Servers.
		s.httpServer = &graceful.Server{
			NoSignalHandling: true,
			Server: &http.Server{
				Addr: s.httpAddr,
			},
			Timeout: 3 * time.Second,
		}

		// Instrumentation.
		http.Handle(s.instrumentation.GetHTTPEndpoint(), s.instrumentation.GetHTTPHandler())

		// HTTP server.
		go func() {
			s.Service().Log().Line("msg", "HTTP server starts to listen on '%s'", s.httpAddr)
			err := s.httpServer.ListenAndServe()
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}()
	})
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "server",
		"type": "service",
	}

	s.bootOnce = sync.Once{}
	s.closer = make(chan struct{}, 1)
	s.shutdownOnce = sync.Once{}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetHTTPAddress(httpAddr string) {
	s.httpAddr = httpAddr
}

func (s *service) SetInstrumentation(i systemspec.Instrumentation) {
	s.instrumentation = i
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.shutdownOnce.Do(func() {
		close(s.closer)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			// Stop the HTTP server gracefully and wait some time for open
			// connections to be closed. Then force it to be stopped.
			s.httpServer.Stop(s.httpServer.Timeout)
			<-s.httpServer.StopChan()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (s *service) Validate() error {
	// Dependencies.

	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	// Settings.

	if s.httpAddr == "" {
		return maskAnyf(invalidConfigError, "HTTP address must not be empty")
	}

	return nil
}
