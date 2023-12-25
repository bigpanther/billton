package app_test

import (
	"github.com/bigpanther/billton/internal/app"
	"github.com/bigpanther/billton/internal/middleware"
	"github.com/gin-gonic/gin"
)

// setupTestRoutes sets up the routes under the given prefix
func setupTestRoutes(prefix string, f func(*gin.RouterGroup)) *gin.Engine {
	engine := gin.Default()
	db := app.InitDb()
	engine.Use(middleware.Database(db))
	engine.MaxMultipartMemory = 8 << 20
	f(engine.Group(prefix))
	return engine
}
