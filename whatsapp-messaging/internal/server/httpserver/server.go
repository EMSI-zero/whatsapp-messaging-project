package httpserver

import (
	"fmt"
	"net/http"
	"os"
	"whatapp-messaging/api/rest"
	"whatapp-messaging/internal/logger"

	"github.com/gin-gonic/gin"
)

var EnvListenAddress = "SRV_LISTEN_ADDRESS"

func StartServer() error {
	engine := gin.Default()
	rest.Addroutes(engine.RouterGroup)

	address, err := GetListenAddressEnv()
	if err != nil {
		return err
	}

	httpserver := &http.Server{Addr: address, Handler: engine}
	go func() {
		logger.Info("http server listening on %s", address)
		if httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panic(err)
		}
	}()

	return nil
}

func GetListenAddressEnv() (string, error) {
	address := os.Getenv(EnvListenAddress)
	if address == "" {
		return "", fmt.Errorf("env %s must be specified", EnvListenAddress)
	}
	return address, nil
}
