package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/config"
	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/pkg/logger"

	"github.com/labstack/echo/v4"

	"github.com/cnrywjd11/go-echo-pkg-oriented-layout/handler"
)

func main() {
	confFilePath := flag.String("conf", "", "config file path")
	logLevel := flag.String("level", "info", "log level")
	flag.Parse()

	logger.SetLogLevel(*logLevel)

	config.InitializeConfig(*confFilePath)

	server := Server{config.GetConfig().Server}
	server.Initialize()
	server.Run()
}

type Server struct {
	config.ServerConfig
}

func (s *Server) Initialize() {
	// Server init
}

func (s *Server) Run() {
	e := echo.New()

	handler.RegisterMiddleware(e)
	handler.RegisterRoutes(e)

	l := logger.JsonLogger()

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", s.Port)); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				l.Infof("Echo Start: %v", err)
			} else {
				l.Fatalf("Echo Start: %v", err)
			}
		}
	}()

	// Wait shutdown signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-shutdown

	timeout := time.Duration(s.ShutdownTimeoutInSecond) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	l.Infof("Shutdown start timeout(%s)", timeout)
	l.Infof("Shutdown server start")
	err := e.Shutdown(ctx)
	l.Infof("Shutdown server finished: %v", err)
}
