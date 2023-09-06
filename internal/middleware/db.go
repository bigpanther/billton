package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop/v6"
)

// Database middleware add database connection to context
func Database(db *pop.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}

}
