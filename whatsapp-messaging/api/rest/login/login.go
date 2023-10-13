package login

import (
	"whatsapp-messaging/internal/logger"

	"github.com/gin-gonic/gin"
)

func AddRoutes(parent *gin.RouterGroup) {
	wa := parent.Group("/whatsapp")

	wa.POST("/login", func(ctx *gin.Context) {
		logger.Info(ctx, "hellloooo")
	})

}
