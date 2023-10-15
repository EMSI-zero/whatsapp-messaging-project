package login

import (
	"encoding/json"
	"net/http"
	"whatsapp-messaging/internal/logger"
	"whatsapp-messaging/services/whatsapp/client"

	"github.com/gin-gonic/gin"
)

func AddRoutes(parent *gin.RouterGroup) {
	wa := parent.Group("/whatsapp")

	wa.POST("/login", func(c *gin.Context) {
		ctx := c.Request.Context()

		msg, qrCode, qrTimeout, err := client.Login(ctx)
		if err != nil {
			c.Error(err)
		}

		if msg == "reconnected successfully" {
			c.JSON(http.StatusAccepted, "reconnected successfully")
		}

		res, err := json.Marshal(&LoginResponse{QRCode: qrCode, QRTimeout: qrTimeout})
		if err != nil {
			logger.Error(ctx, err)
			c.Error(err)
		}

		c.JSON(http.StatusOK, res)
	})

}
