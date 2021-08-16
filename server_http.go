package jinx

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	WebPort    int
	WebTimeOut int
	ServeOpts  struct {
		Port    WebPort
		TimeOut WebTimeOut
	}
)

type Server struct {
	errChan    chan error
	httpServer *http.Server
	Port       WebPort
	TimeOut    WebTimeOut
}

func opening() string {
	return `
-------------------------------
      WELCOME TO WILDRIFT
-------------------------------
LISTENING ON PORT : %d
...............................
`
}

func closing() string {
	return `
...............................
       server stopping
`
}

//Run function receive http.Handler to execute.
func (s *Server) Run(handler http.Handler) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.httpServer.Handler = handler
	// Description µ micro service
	fmt.Println(fmt.Sprintf(
		opening(),
		s.Port,
	))
	go func() {
		s.errChan <- s.httpServer.ListenAndServe()
	}()
	s.waitForSignals(ctx)
	return
}

//NewHTTPServer is function to return *http.server with parameter included address port,
//read timeout and write timeout.
func NewHTTPServer(opts ServeOpts) *Server {
	return &Server{httpServer: &http.Server{
		Addr:         fmt.Sprintf(":%d", opts.Port),
		ReadTimeout:  time.Duration(opts.TimeOut) * time.Second,
		WriteTimeout: time.Duration(opts.TimeOut) * time.Second,
	}, Port: opts.Port}
}

func (s *Server) waitForSignals(ctx context.Context) {
	// Do not make the application hang when it is shutdown.
	ctxOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-interrupt:
			fmt.Println(" <-- ups, server was interrupt")
			s.Stop(ctxOut)
			return
		case err := <-s.errChan:
			fmt.Println(" oh, server was error", err)
			s.Stop(ctxOut)
			return
		}
	}
}

// Stop function for stopping server gracefully.
func (s *Server) Stop(ctx context.Context) {
	fmt.Println(closing())
	if err := s.httpServer.Shutdown(ctx); err != nil {
		if err := s.httpServer.Close(); err != nil {
			fmt.Println("(!) with error :", err)
		}
	}
}
