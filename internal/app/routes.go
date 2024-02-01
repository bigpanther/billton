package app

import (
	"github.com/bigpanther/billton/internal/firebase"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, firebase firebase.Firebase) {
	// Group all routes under v1
	v1 := r.Group("v1")
	warranties := v1.Group("warranties")
	WarrantiesRoutes(warranties)
	users := v1.Group("users")
	UserRoutes(users)
	self := v1.Group("self")
	SelfRoutes(self, firebase)
}
