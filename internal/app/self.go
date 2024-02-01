package app

import (
	"net/http"

	"log/slog"

	"github.com/bigpanther/billton/internal/firebase"
	"github.com/gin-gonic/gin"
)

func SelfRoutes(rg *gin.RouterGroup, firebase firebase.Firebase) {
	rg.GET("/", selfGet)
	rg.POST("/device-register", selfPostDeviceRegister(firebase))
	rg.POST("/device-remove", selfPostDeviceRemove(firebase))
}

func selfGet(c *gin.Context) {
	c.JSON(http.StatusOK, loggedInUser(c))
}

func selfPostDeviceRegister(f firebase.Firebase) func(c *gin.Context) {
	return func(c *gin.Context) {

		deviceB := deviceID{}
		if err := c.Bind(&deviceB); err != nil {
			slog.Error("error binding deviceid:", err)
			_ = c.Error(err)
			return
		}
		if err := f.SubscribeToTopics(c, loggedInUser(c), deviceB.Token); err != nil {
			_ = c.Error(err)
			return
		}
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}

func selfPostDeviceRemove(f firebase.Firebase) func(c *gin.Context) {
	return func(c *gin.Context) {
		deviceB := deviceID{}
		if err := c.Bind(&deviceB); err != nil {
			slog.Error("error binding deviceid:", err)
			_ = c.Error(err)
			return
		}
		if err := f.UnSubscribeToTopics(c, loggedInUser(c), deviceB.Token); err != nil {
			_ = c.Error(err)
			return
		}
		c.Writer.WriteHeader(http.StatusNoContent)
	}
}

type deviceID struct {
	Token string `json:"token"`
}
