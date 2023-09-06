package context

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop/v6"
)

type AppContext struct {
	*gin.Context
	DB *pop.Connection
}
