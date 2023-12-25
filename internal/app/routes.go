package app

import (
	"github.com/bigpanther/billton/internal/firebase"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, firebase firebase.Firebase) {
	warranties := r.Group("warranties")
	WarrantiesRoutes(warranties)
	users := r.Group("users")
	UserRoutes(users)
	self := r.Group("self")
	SelfRoutes(self, firebase)
	chains := []gin.HandlersChain{
		warranties.Handlers,
		users.Handlers,
		self.Handlers,
	}
	// Group all routes under v1
	v1 := r.Group("v1")
	for _, chain := range chains {
		v1.Handlers = append(v1.Handlers, chain...)
	}

}
