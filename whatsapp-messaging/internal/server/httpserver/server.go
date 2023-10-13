package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"whatsapp-messaging/api/rest"
	"whatsapp-messaging/internal/logger"

	"github.com/gin-gonic/gin"
)

var EnvListenAddress = "SRV_LISTEN_ADDRESS"

func StartServer() error {
	engine := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	rest.AddRoutes(&engine.RouterGroup)


	address, err := GetListenAddressEnv()
	if err != nil {
		return err
	}

	httpserver := &http.Server{Addr: address, Handler: engine}
	go func() {
		logger.Info(context.Background(), "http server listening on %s", address)
		if httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panic(context.Background(), err)
		}
	}()

	srv := httpserver
	chInterrupt := make(chan os.Signal, 1)
	signal.Notify(chInterrupt, os.Interrupt)
	chTerm := make(chan os.Signal, 1)
	signal.Notify(chTerm, syscall.SIGTERM)
	select {
	case <-chInterrupt:
		logger.Info(context.Background(), "server interrupted")
		stopServer(srv)
	case <-chTerm:
		logger.Info(context.Background(), "server received SIGTERM")
	}
	logger.Info(context.Background(), "done")

	return nil
}

func stopServer(srv *http.Server) {
	partStopped := make(chan struct{}, 20)
	go func() {
		const timeout = 5 * time.Second
		time.Sleep(timeout)
		partStopped <- struct{}{}
	}()
	go func() {
		logger.Info(context.Background(), "stopping http server...")
		srvError := srv.Close()
		if srvError != nil {
			logger.Info(context.Background(), "error: http server termination failed: %v", srvError)
		} else {
			logger.Info(context.Background(), "stopped http server")
		}
		partStopped <- struct{}{}
	}()
	<-partStopped
}

func GetListenAddressEnv() (string, error) {
	address := os.Getenv(EnvListenAddress)
	if address == "" {
		return "", fmt.Errorf("env %s must be specified", EnvListenAddress)
	}
	return address, nil
}
